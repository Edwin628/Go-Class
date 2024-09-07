package api

import (
	"go-class/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建房间
func CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := models.CreateRoom(room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room created successfully"})
}

// 获取房间列表
func GetRooms(c *gin.Context) {
	rooms := models.GetAllRooms()
	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}
