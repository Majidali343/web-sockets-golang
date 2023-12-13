package main

import (
	"fmt"

	chating "github.com/Majidali343/web-sockets-golang/internal/handlers"

	tokens "github.com/Majidali343/web-sockets-golang/internal/Middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ws", chating.HandleConnections)
	// r.GET("/roomws", chating.HandleroomConnections)

	go chating.HandleMessages()
	// go chating.HandleroomMessages()

	r.POST("/createroom", tokens.AuthMiddleware(), chating.CreateRom)
	r.GET("/getroom", chating.GetRoomsByUserEmail)
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
