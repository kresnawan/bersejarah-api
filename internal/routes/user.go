package routes

import (
	"github.com/gin-gonic/gin"
)

func handleUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{"message": "Helo! " + name})
}

func UsersRoutes(rg *gin.RouterGroup) {
	UsersRoute := rg.Group("/users")
	{
		UsersRoute.GET("/:name", handleUser)
	}
}
