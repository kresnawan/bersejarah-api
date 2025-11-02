package routes

import (
	"app/internal/handler"

	"github.com/julienschmidt/httprouter"
)

var Router = httprouter.New()

func UserRoute() {
	Router.GET("/", handler.IndexPage)
	Router.GET("/hello/:name", handler.HelloPage)
}
