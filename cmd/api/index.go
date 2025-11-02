package main

import (
	"app/internal/routes"
	"app/internal/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {

	routes.UserRoute()

	fmt.Printf("Server starting on port 3030...\n\n")
	err := http.ListenAndServe(":3030", routes.Router)

	if err != nil {
		log.Fatalf("\nServer failed to start: %v", err)
	}

	defer storage.Db.Close()
}
