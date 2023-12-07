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

	r.GET("/", chating.HomePage)
	r.GET("/ws", chating.HandleConnections)

	go chating.HandleMessages()

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
