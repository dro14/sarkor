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

type Login struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {

	user, err := validateUserInput(c)
	if err != nil {
		log.Println("user input is invalid:", err)
		return
	}

	result := database.DB.Create(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func AuthenticateUser(c *gin.Context) {

	var login Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("login input is invalid:", err)
		return
	}

	var user models.User
	result := database.DB.Where("login = ?", login).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login"})
		log.Println("invalid login:", login)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		log.Println("invalid password for ", login)
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

func GetUserByName(c *gin.Context) {

	name := c.Param("name")

	var user models.User
	result := database.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
		"age":  user.Age,
	})
}
