package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	api := router.Group("/api")
	{
		api.POST("/sum", Sum)
	}

	// Start and run the server
	router.Run(":8000")
}

type SumRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func Sum(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  "Unable to get request",
		})
		return
	}

	sumRequest := SumRequest{}
	err = json.Unmarshal(body, &sumRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  "Cannot unmarshal body",
		})
		return
	}

	a, err := strconv.Atoi(sumRequest.A)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "a is not a valid number or null",
		})
		return
	}

	b, err := strconv.Atoi(sumRequest.B)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "b is not a valid number or null",
		})
		return
	}

	sum := a + b

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"a":     a,
		"b":     b,
		"hasil": sum,
	})
}
