package main

import (
	"github.com/Kintenis/go-api/database"
	"github.com/Kintenis/go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
