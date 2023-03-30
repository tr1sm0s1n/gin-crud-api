package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.Run("localhost:8080")
}
