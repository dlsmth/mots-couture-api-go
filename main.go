package main

import (
	"github.com/gin-gonic/gin"
	"mc.com/rest-api/db"
	"mc.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
