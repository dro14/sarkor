package middleware

import (
	"net/http"

	"github.com/dro14/sarkor/utils"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	cookie, err := c.Request.Cookie("SESSTOKEN")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
		c.Abort()
		return
	}

	token := cookie.Value
	claims, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("userID", claims.UserID)
	c.Next()
}
