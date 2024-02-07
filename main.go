package main

import (
	"github.com/gin-gonic/gin"
	"mario-mtz.com/rest-api/db"
	"mario-mtz.com/rest-api/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
