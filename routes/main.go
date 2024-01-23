package routes

import (
	"main/handlers/excelhandler"

	"github.com/Elhjay/qrVra/main/handlers/excelhandler"
	"github.com/Elhjay/qrVra/main/handlers/vcardhandler"

	"github.com/gin-gonic/gin"
)

func MainRoutes(server *gin.Engine) {

	server.POST("/excel", func(ctx *gin.Context) {
		excelhandler.ExcelFileHandler(ctx)
	})
	server.GET("vcf/:id", func(ctx *gin.Context) {
		vcardhandler.VCFHandler(ctx)
	})

	server.PATCH("/update/:id", func(ctx *gin.Context) {
		vcardhandler.UpdateVCardInfoHandler(ctx)
	})
}
