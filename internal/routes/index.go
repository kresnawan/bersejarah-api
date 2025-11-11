package routes

import (

	// "fmt"
	// "log"
	// "net/http"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func RouterInit() {
	var Router = gin.Default()
	Router.SetTrustedProxies([]string{"127.0.0.1"})

	Router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		fmt.Printf("New user connected: %s\n", conn.LocalAddr().String())

		defer conn.Close()

		for {
			// Read message from client
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}
			fmt.Printf("Received message: %s\n", message)

			// Echo message back to client
			// err = conn.WriteMessage(websocket.TextMessage, message)
			// if err != nil {
			// 	fmt.Println("Error writing message:", err)
			// 	break
			// }
		}
	})

	MainRouter := Router.Group("/api/v1")
	{
		UsersRoutes(MainRouter)
		TestRoutes(MainRouter)
		AuthRoutes(MainRouter)
		DataTempatRoutes(MainRouter)
	}
	Router.Run()
}
