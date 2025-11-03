package main

import (
	// "app/internal/routes"
	"app/internal/routes"
	"app/internal/storage"
	// "fmt"
	// "log"
	// "net/http"
)

func main() {

	routes.RouterInit()
	defer storage.Db.Close()
}
