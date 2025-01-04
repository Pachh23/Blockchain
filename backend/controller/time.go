package controller

import (
	"net/http"

	"blockchain.com/bc-67/config"
	"blockchain.com/bc-67/entity"
	"github.com/gin-gonic/gin"
)

// GetAllRooms - ดึงข้อมูลห้องทั้งหมด
func GetAllTimes(c *gin.Context) {
	db := config.DB()
	var times []entity.Times
	if err := db.Find(&times).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูลได้"})
		return
	}
	c.JSON(http.StatusOK, times)
}
