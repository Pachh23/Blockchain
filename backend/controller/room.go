package controller

import (
	"net/http"

	"blockchain.com/bc-67/config"
	"blockchain.com/bc-67/entity"
	"github.com/gin-gonic/gin"
)

// GetAllRooms - ดึงข้อมูลห้องทั้งหมด
func GetAllRooms(c *gin.Context) {
	db := config.DB()
	var rooms []entity.Room
	if err := db.Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูลได้"})
		return
	}
	c.JSON(http.StatusOK, rooms)
}
