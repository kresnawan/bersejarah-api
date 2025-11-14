package routes

import (

	// "fmt"
	// "log"
	// "net/http"

	// "app/internal/middleware"

	"app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RouterInit() {
	// gin.SetMode(gin.ReleaseMode)

	Router := gin.Default()
	Router.Use(middleware.CORSMiddleware())
	Router.SetTrustedProxies([]string{"127.0.0.1"})

	MainRouter := Router.Group("/api/v1")
	{
		UsersRoutes(MainRouter)
		TestRoutes(MainRouter)
		AuthRoutes(MainRouter)
		DataTempatRoutes(MainRouter)
	}

	Router.Run()
}
