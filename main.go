package main

import (
	"net/http"

	"github.com/Ali-Assar/go-REST-Api.git/db"
	"github.com/Ali-Assar/go-REST-Api.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id")
	server.POST("/events", createEvent)

	server.Run(":8080") //listening on port 8080

}
