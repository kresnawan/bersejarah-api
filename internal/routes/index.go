package routes

import (

	// "fmt"
	// "log"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	var Router = gin.Default()

	MainRouter := Router.Group("/api/v1")
	{
		UsersRoutes(MainRouter)
	}
	Router.Run()
}
