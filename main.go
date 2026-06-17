package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tr1sm0s1n/gin-crud-api/docs"
)

// @Summary		Create new certificate
// @Tags 		Protected
// @Security	Auth
// @Accept  	json
// @Produce		json
// @Param		certificate body Certificate true "Certificate Data"
// @Success 	201 {object} Certificate
// @Failure 	400 {object} map[string]any
// @Failure 	401 {object} map[string]any
// @Router		/create [post]
func create(c *gin.Context) {
	var newCertificate Certificate
	if err := c.BindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := getIndex(newCertificate.ID); exists {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Already Exists"})
		return
	}

	certificates = append(certificates, newCertificate)
	c.IndentedJSON(http.StatusCreated, newCertificate)
}

// @Summary		Read all certificates
// @Tags 		Protected
// @Security	Auth
// @Produce		json
// @Success 	200 {object} []Certificate
// @Failure 	401 {object} map[string]any
// @Router 		/read [get]
func readAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, certificates)
}

// @Summary		Read a single certificate data corresponding to an ID
// @Tags 		Protected
// @Security	Auth
// @Produce		json
// @Param 		id path string true "Certificate ID"
// @Success 	200 {object} Certificate
// @Failure 	401 {object} map[string]any
// @Failure 	404 {object} map[string]any
// @Router 		/read/{id} [get]
func readOne(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if i, exists := getIndex(id); exists {
		c.IndentedJSON(http.StatusOK, certificates[i])
		return
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}

// @Summary		Update or upsert a certificate
// @Tags 		Protected
// @Security	Auth
// @Accept  	json
// @Produce		json
// @Param 		id path string true "Certificate ID"
// @Param		property body Certificate true "Certificate Data"
// @Success 	200 {object} Certificate
// @Success 	201 {object} Certificate
// @Failure 	400 {object} map[string]any
// @Failure 	401 {object} map[string]any
// @Router		/update/{id} [put]
func update(c *gin.Context) {
	var updated Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&updated); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if i, exists := getIndex(id); exists {
		certificates[i] = updated
		c.IndentedJSON(http.StatusOK, updated)
		return
	}

	certificates = append(certificates, updated)
	c.IndentedJSON(http.StatusCreated, updated)
}

// @Summary		Delete an existing certificate
// @Tags 		Protected
// @Security	Auth
// @Produce		json
// @Param 		id path string true "Certificate ID"
// @Success 	200 {object} Certificate
// @Failure 	400 {object} map[string]any
// @Failure 	401 {object} map[string]any
// @Failure 	404 {object} map[string]any
// @Router		/delete/{id} [delete]
func delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if i, exists := getIndex(id); exists {
		deleted := certificates[i]
		certificates = append(certificates[:i], certificates[i+1:]...)
		c.IndentedJSON(http.StatusOK, deleted)
		return
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
}

// @title						Gin-CRUD-API
// @version						1.0
// @host						localhost:8081
// @schemes						http
// @tag.name 					Protected
// @tag.description 			X-API-KEY is required.
// @securityDefinitions.apikey	Auth
// @name 						X-API-KEY
// @in 							header
func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(authority())
	router.POST("/create", create)
	router.GET("/read", readAll)
	router.GET("/read/:id", readOne)
	router.PUT("/update/:id", update)
	router.DELETE("/delete/:id", delete)
	router.Run(":8081")
}
