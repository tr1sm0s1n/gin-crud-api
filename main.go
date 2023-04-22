package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type certificate struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
	Grade  string `json:"grade"`
	Date   string `json:"date"`
}

var certificates = []certificate{
	{ID: 36, Name: "Abigail", Course: "CED", Grade: "A", Date: "20-1-2044"},
	{ID: 39, Name: "Leopold", Course: "CBA", Grade: "S", Date: "20-1-2044"},
	{ID: 42, Name: "Zachary", Course: "CHF", Grade: "B", Date: "20-1-2044"},
}

func readAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, certificates)
}

func readOne(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	certificate, err := getOne(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, certificate)
}

func create(c *gin.Context) {
	var newCertificate certificate

	if err := c.BindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	certificates = append(certificates, newCertificate)
	c.IndentedJSON(http.StatusCreated, newCertificate)
}

func update(c *gin.Context) {
	var updated certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err := c.BindJSON(&updated); err != nil {
		return
	}

	certificate, err := updateOne(id, updated)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, certificate)
}

func delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	index, err := getIndex(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	deleted := certificates[*index]
	certificates = append(certificates[:*index], certificates[*index+1:]...)
	c.IndentedJSON(http.StatusOK, deleted)
}

func main() {
	cl, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	chainID, err := cl.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	networkID, err := cl.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	
	printNID := fmt.Sprintf("Network ID: %d", networkID)
	printCID := fmt.Sprintf("Chain ID: %d", chainID)
	fmt.Println(printNID)
	fmt.Println(printCID)

	router := gin.Default()
	router.GET("/read", readAll)
	router.GET("/read/:id", readOne)
	router.POST("/create", create)
	router.PUT("/update/:id", update)
	router.DELETE("/delete/:id", delete)
	router.Run("localhost:8080")
}
