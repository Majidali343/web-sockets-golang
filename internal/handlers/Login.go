package chat

import (
	"net/http"

	dbconnect "github.com/Majidali343/web-sockets-golang/internal/db"
	"github.com/gin-gonic/gin"

	tokens "github.com/Majidali343/web-sockets-golang/internal/Middleware"
)

func Login(c *gin.Context) {
	var user MyUser

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var dbUser User

	// Connect to the database
	db := dbconnect.Dbconnection()
	defer db.Close()

	// Check if the user with the provided email exists
	if err := db.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify the password
	if dbUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, err := tokens.GenerateAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	refreshToken, err := tokens.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	LogedUser.Useremail = dbUser.Email

	// Respond with user ID and token
	c.JSON(http.StatusOK, gin.H{"user_email": dbUser.Email, "access_token": accessToken,
		"refresh_token": refreshToken})
}
