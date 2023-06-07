package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dro14/sarkor/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func validateUserInput(c *gin.Context) (*models.User, error) {

	login := c.PostForm("login")
	if len(login) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login is required"})
		return nil, errors.New("login is required")
	}

	password := c.PostForm("password")
	if len(password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return nil, errors.New("password is required")
	}

	name := c.PostForm("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return nil, errors.New("name is required")
	}

	ageStr := c.PostForm("age")
	if len(ageStr) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Age is required"})
		return nil, errors.New("age is required")
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil || age < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid age"})
		return nil, errors.New("invalid age: " + ageStr)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return nil, errors.New("can't hash password: " + err.Error())
	}

	user := &models.User{
		Login:    login,
		Password: string(hashedPassword),
		Name:     name,
		Age:      age,
	}

	return user, nil
}
