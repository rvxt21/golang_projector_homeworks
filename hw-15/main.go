package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	server := NewToursServer()
	router.GET("/tours/", server.GetToursHandler)
	router.POST("/tours/", server.CreateTourHandler)
}
