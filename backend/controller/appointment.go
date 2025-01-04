package controller

import (
	"encoding/json"
	"net/http"

	"blockchain.com/bc-67/blockchain" // เพิ่มการ import package blockchain
	"blockchain.com/bc-67/config"
	"blockchain.com/bc-67/entity"
	"github.com/gin-gonic/gin"
)

var blockchainInstance = blockchain.NewBlockchain() // สร้าง instance ของ Blockchain

func CreateAppointment(c *gin.Context) {
	var appointment entity.Appointment

	// Bind JSON data to appointment struct
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// Create a new appointment entry
	if err := db.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	// Convert the appointment object to JSON for storing in the blockchain
	appointmentData, _ := json.Marshal(appointment)

	// Create a new block using the appointment data
	newBlock, err := blockchainInstance.CreateBlock(string(appointmentData)) // รับค่าคืนมาทั้ง Block และ Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create block"})
		return
	}

	// ส่งข้อมูลตอบกลับรวมทั้งข้อมูล appointment และ block
	c.JSON(http.StatusCreated, gin.H{
		"message": "Appointment created successfully",
		"data":    appointment,
		"block":   newBlock,
	})
}

func GetAllBlocks(c *gin.Context) {
	c.JSON(http.StatusOK, blockchainInstance.Chain)
}

/*
func GetAllBlocks(c *gin.Context) {
	// ดึงข้อมูลจาก Blockchain
	blocks := blockchainInstance.Chain

	// ดึงข้อมูล Appointment จากฐานข้อมูล
	var appointments []entity.Appointment
	db := config.DB()
	results := db.
		Preload("Department").
		Preload("Time").
		Preload("Room").
		Find(&appointments)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	// สร้างผลลัพธ์ที่รวมข้อมูลจาก Blockchain และ Appointment
	var combinedResult []map[string]interface{}
	for _, block := range blocks {
		// ตรวจสอบว่า Block Data เกี่ยวข้องกับ Appointment หรือไม่
		for _, appointment := range appointments {
			// คุณสามารถกำหนดเงื่อนไขการจับคู่ได้ที่นี่ (เช่น การตรวจสอบ `block.Data` กับ `appointment.ID`)
			if block.Data == string(appointment.ID) { // สมมุติว่า block.Data เก็บข้อมูลที่เกี่ยวข้องกับ Appointment.ID
				// ผสมข้อมูลจาก Blockchain กับ Appointment
				combinedResult = append(combinedResult, map[string]interface{}{
					"block":       block,
					"appointment": appointment,
				})
			}
		}
	}

	// ส่งผลลัพธ์ทั้งหมดกลับไปยังผู้ใช้งาน
	c.JSON(http.StatusOK, combinedResult)
}
*/
/*
// GetAllAppointments - GET /appointments

	func GetAllAppointments(c *gin.Context) {
		var appointments []entity.Appointment
		db := config.DB()
		// Fetch all appointments and preload the associated patient and department data
		results := db.Preload("Department").Preload("Time").Preload("Department.Name").Find(&appointments)
		if results.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, appointments)
	}
*/
func GetAllAppointments(c *gin.Context) {
	var appointments []entity.Appointment
	db := config.DB()

	// Preload Department และ Rooms
	results := db.
		Preload("Department"). // โหลด Rooms ที่เกี่ยวข้องกับ Department
		Preload("Time").
		Preload("Room"). // โหลด Time
		Find(&appointments)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

// GetAppointmentByID - GET /appointments/:id
func GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	var appointment entity.Appointment

	db := config.DB()

	// Fetch appointment by ID and preload patient and department data
	if err := db.Preload("Patient").Preload("Department").First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Return appointment data with preloaded patient and department data
	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment fetched successfully",
		"data":    appointment,
	})
}

// UpdateAppointment - PUT /appointments/:id
func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment entity.Appointment

	db := config.DB()

	// Find the appointment by ID
	if err := db.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Bind the updated data to the appointment struct
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated appointment data
	if err := db.Save(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment updated successfully",
		"data":    appointment,
	})
}

// DeleteAppointment - DELETE /appointments/:id
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()

	// Delete the appointment by ID
	if err := db.Delete(&entity.Appointment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment deleted successfully",
	})
}
