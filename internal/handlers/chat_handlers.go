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
	RoomName string `json:"roomname"`
}

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request

	Useremail := c.Query("email")
	RoomName := c.Query("roomname")

	fmt.Println(Useremail)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if Useremail != "" {
		clients[conn] = Useremail

	} else {
		clients[conn] = RoomName
	}

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
			} else if clientUsername == msg.RoomName {

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
