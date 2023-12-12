package chat

import (
	"net/http"

	dbconnect "github.com/Majidali343/web-sockets-golang/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Room struct {
	RoomName string `form:"roomname"`
}

type ChatRoom struct {
	gorm.Model
	UserEmail string `form:"useremail"`
	Roomname  string `form:"roomname"`
}

func CreateRom(c *gin.Context) {

	db := dbconnect.Dbconnection()
	defer db.Close()

	db.AutoMigrate(&ChatRoom{})

	var room Room

	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	RoomCreator := LogedUser.Useremail

	var newRoom ChatRoom
	newRoom.UserEmail = RoomCreator
	newRoom.Roomname = room.RoomName

	if err := db.Create(&newRoom).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room Created successfully successfully"})

}

func GetRoomsByUserEmail(c *gin.Context) {
	db := dbconnect.Dbconnection()
	defer db.Close()

	userEmail := LogedUser.Useremail

	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User email is required"})
		return
	}

	var rooms []ChatRoom

	// Query the database to get rooms associated with the user's email
	if err := db.Where("user_email = ?", userEmail).Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rooms", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}
