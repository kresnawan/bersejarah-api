package routes

import (

	// "fmt"
	// "log"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	var Router = gin.Default()
	Router.SetTrustedProxies([]string{"127.0.0.1"})

	MainRouter := Router.Group("/api/v1")
	{
		UsersRoutes(MainRouter)
		TestRoutes(MainRouter)
	}
	Router.Run()
}
