package main

import (
	"main/database"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//vcardService.AddProductService()

	server := gin.Default()

	database.ConnectDB()
	routes.MainRoutes(server)
	server.Run(":8080")

}
