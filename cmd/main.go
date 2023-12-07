package main

import (
	"fmt"
	"net/http"

	chating "github.com/Majidali343/web-sockets-golang/internal/handlers"
	"github.com/gorilla/websocket"

	tokens "github.com/Majidali343/web-sockets-golang/internal/Middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", homePage)
	r.GET("/ws", handleConnections)

	go handleMessages()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})

	r.POST("/login", chating.Login)
	r.POST("/register", chating.Register)
	r.POST("/refresh", tokens.RefreshTokenHandler)

	return r
}

func main() {
	router := SetupRouter()
	router.Run(":8080")

	fmt.Println("Server started on :8080")

}

func homePage(c *gin.Context) {
	w := c.Writer
	fmt.Fprintf(w, "Chat room is working!")
}

func handleConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			return
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
