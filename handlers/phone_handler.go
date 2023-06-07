package handlers

import (
	"log"
	"net/http"

	"github.com/dro14/sarkor/database"
	"github.com/dro14/sarkor/models"
	"github.com/gin-gonic/gin"
)

// AddPhone is a handler that adds a phone number to the database
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "You already have a phone number, please update it"})
		log.Println("failed to add phone:", result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phone added successfully"})
}

// GetPhone is a handler that retrieves a phone number from the database
func GetPhone(c *gin.Context) {

	number := c.Query("q")

	var phones []models.Phone
	result := database.DB.Where("phone = ?", number).Find(&phones)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve phones"})
		log.Println("failed to retrieve phones:", result.Error)
		return
	}

	c.JSON(http.StatusOK, phones)
}

// UpdatePhone is a handler that updates a phone number in the database
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

// DeletePhone is a handler that deletes a phone number from the database
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
