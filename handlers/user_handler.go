package handlers

import (
	"log"
	"net/http"

	"github.com/dro14/sarkor/database"
	"github.com/dro14/sarkor/models"
	"github.com/dro14/sarkor/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Credentials is a struct for login input
type Credentials struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser is a handler for registering a new user
func RegisterUser(c *gin.Context) {

	user, err := validateUserInput(c)
	if err != nil {
		log.Println("user input is invalid:", err)
		return
	}

	result := database.DB.Create(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User was already registered"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// AuthenticateUser is a handler for user authentication
func AuthenticateUser(c *gin.Context) {

	var credentials Credentials
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("login input is invalid:", err)
		return
	}

	var user models.User
	result := database.DB.Where("login = ?", credentials.Login).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login"})
		log.Println("invalid login:", credentials)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		log.Println("invalid password for ", credentials)
		return
	}

	claims := &utils.Claims{
		Login:  user.Login,
		UserID: user.ID,
	}

	token, err := utils.GenerateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		log.Println("can't generate token:", err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "SESSTOKEN",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful"})
}

// GetUserByName is a handler for getting user by name
func GetUserByName(c *gin.Context) {

	name := c.Param("name")

	var user models.User
	result := database.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Login = ""
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
