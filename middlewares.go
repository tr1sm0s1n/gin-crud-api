package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const expectedKey = "XmoKQLLRC5M3ufz8Y68jXti8bafd9MAn"

func authority() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid X-API-KEY"})
			return
		}

		c.Next()
	}
}