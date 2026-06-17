package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/create", create)

	newCertificate := certificates[0]
	newCertificate.ID = 101
	
	jsonValue, _ := json.Marshal(newCertificate)
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestReadAllHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/read", readAll)

	req, _ := http.NewRequest("GET", "/read", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var certificates []Certificate
	json.Unmarshal(w.Body.Bytes(), &certificates)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, certificates)
}

func TestReadOneHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/read/:id", readOne)

	param := strconv.Itoa(1)
	req, _ := http.NewRequest("GET", "/read/"+param, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var certificate Certificate
	json.Unmarshal(w.Body.Bytes(), &certificate)
	fmt.Println(certificate.Name)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, certificate.Name, "John Doe")
	assert.Equal(t, certificate.Course, "Advanced React")
}

func TestUpdateHandler(t *testing.T) {
	r := SetUpRouter()
	r.PUT("/update/:id", update)

	updated := certificates[0]
	updated.Name = "Delta Four"

	param := strconv.Itoa(updated.ID)
	jsonValue, _ := json.Marshal(updated)
	req, _ := http.NewRequest("PUT", "/update/"+param, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteHandler(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/delete/:id", delete)

	param := strconv.Itoa(2)
	req, _ := http.NewRequest("DELETE", "/delete/"+param, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
