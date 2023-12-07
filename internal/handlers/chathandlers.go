package chat

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HomePage(c *gin.Context) {
	w := c.Writer
	fmt.Fprintf(w, "Chat room is working!")
}

func HandleConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request

	Useremail := c.Query("email")

	fmt.Println(Useremail)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = Useremail

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

func HandleMessages() {
	for {
		msg := <-broadcast

		for client, clientUsername := range clients {
			fmt.Println(clientUsername)
			fmt.Println(msg.Receiver)
			if clientUsername == msg.Receiver {
				err := client.WriteJSON(msg)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
