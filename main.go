package main

import (
	"main/routes"

	"main/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//vcardService.AddProductService()

	server := gin.Default()

	database.ConnectDB()
	routes.MainRoutes(server)

	//lambda.Start(server.Handle)
	server.Run(":8080")

}
