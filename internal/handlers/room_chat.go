package chat

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Roomclients = make(map[*websocket.Conn]string)
var Roombroadcast = make(chan Message)

func HandleroomConnections(c *gin.Context) {
	w := c.Writer
	r := c.Request

	Roomname := c.Query("roomname")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = Roomname

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

func HandleroomMessages() {
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
