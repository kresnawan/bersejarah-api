package routes

import (
	"app/internal/handler"

	"github.com/gin-gonic/gin"
)

func DataTempatRoutes(rg *gin.RouterGroup) {
	UsersRoute := rg.Group("/data")
	{
		UsersRoute.POST("/", handler.AddDataTempat)
		UsersRoute.POST("/upload", handler.UploadFoto)
		UsersRoute.GET("/", handler.GetAllDataTempat)
	}
}
