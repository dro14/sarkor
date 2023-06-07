package handlers

import "github.com/gin-gonic/gin"

func getUserIDFromToken(c *gin.Context) uint {
	userID, _ := c.Get("userID")
	return userID.(uint)
}
