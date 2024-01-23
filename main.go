package main

import (
	"github.com/Elhjay/qrVra/main/database"
	"github.com/Elhjay/qrVra/main/routes"

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
