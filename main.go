package main

import (
	"github.com/Ali-Assar/go-REST-Api.git/db"
	"github.com/Ali-Assar/go-REST-Api.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRouts(server)

	server.Run(":8080") //listening on port 8080

}
