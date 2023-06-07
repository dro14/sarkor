package handlers

import "github.com/gin-gonic/gin"

// getUserIDFromToken is a helper function that retrieves the user ID from the token
func getUserIDFromToken(c *gin.Context) uint {
	userID, _ := c.Get("userID")
	return userID.(uint)
}
