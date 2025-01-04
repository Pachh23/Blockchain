package controller

import (
	"net/http"

	"blockchain.com/bc-67/config"
	"blockchain.com/bc-67/entity"
	"github.com/gin-gonic/gin"
)

func GetAllDepartments(c *gin.Context) {
	db := config.DB()
	var departments []entity.Departments
	db.Find(&departments)
	c.JSON(http.StatusOK, &departments)
}
