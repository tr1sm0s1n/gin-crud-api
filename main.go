package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var certificates = []Certificate{
	{Id: 36, Name: "Abigail", Course: "CED", Grade: "A", Date: "20-1-2044"},
	{Id: 39, Name: "Leopold", Course: "CBA", Grade: "S", Date: "20-1-2044"},
	{Id: 42, Name: "Zachary", Course: "CHF", Grade: "B", Date: "20-1-2044"},
}

func create(c *gin.Context) {
	var newCertificate Certificate
	if err := c.BindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	i, exists := getIndex(newCertificate.Id)
	if exists {
		certificates[i] = newCertificate
		c.IndentedJSON(http.StatusCreated, newCertificate)
		return
	}

	certificates = append(certificates, newCertificate)
	c.IndentedJSON(http.StatusCreated, newCertificate)
}

func readAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, certificates)
}

func readOne(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	i, exists := getIndex(id)
	if exists {
		c.IndentedJSON(http.StatusOK, certificates[i])
		return
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func update(c *gin.Context) {
	var updated Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if err := c.BindJSON(&updated); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	i, exists := getIndex(id)
	if exists {
		certificates[i] = updated
		c.IndentedJSON(http.StatusOK, updated)
		return
	}
	
	certificates = append(certificates, updated)
	c.IndentedJSON(http.StatusCreated, updated)
}

func delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	i, exists := getIndex(id)
	if exists {
		deleted := certificates[i]
		certificates = append(certificates[:i], certificates[i+1:]...)
		c.IndentedJSON(http.StatusOK, deleted)
		return
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
}

func main() {
	router := gin.Default()
	router.POST("/create", create)
	router.GET("/read", readAll)
	router.GET("/read/:id", readOne)
	router.PUT("/update/:id", update)
	router.DELETE("/delete/:id", delete)
	router.Run("localhost:8080")
}
