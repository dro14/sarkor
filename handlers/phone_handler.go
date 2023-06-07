package handlers

import (
	"log"
	"net/http"

	"github.com/dro14/sarkor/database"
	"github.com/dro14/sarkor/models"
	"github.com/gin-gonic/gin"
)

func AddPhone(c *gin.Context) {

	var phone models.Phone
	err := c.ShouldBindJSON(&phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone data"})
		log.Println("phone input is invalid:", err)
		return
	}

	userID := getUserIDFromToken(c)
	phone.UserID = userID

	result := database.DB.Create(&phone)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add phone"})
		log.Println("failed to add phone:", result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phone added successfully"})
}

func GetPhone(c *gin.Context) {

	number := c.Query("q")

	var phone models.Phone
	result := database.DB.Where("phone_number = ?", number).First(&phone)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve phone"})
		log.Println("failed to retrieve phone:", result.Error)
		return
	}

	c.JSON(http.StatusOK, phone)
}

func UpdatePhone(c *gin.Context) {

	var phone models.Phone
	err := c.ShouldBindJSON(&phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone data"})
		log.Println("phone input is invalid:", err)
		return
	}

	userID := getUserIDFromToken(c)
	phone.UserID = userID

	result := database.DB.Save(&phone)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update phone"})
		log.Println("failed to update phone:", result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phone updated successfully"})
}

func DeletePhone(c *gin.Context) {

	phoneID := c.Param("phone_id")

	var phone models.Phone
	result := database.DB.Where("id = ?", phoneID).Delete(&phone)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete phone"})
		log.Println("failed to delete phone:", result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phone deleted successfully"})
}
