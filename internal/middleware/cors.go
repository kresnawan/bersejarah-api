package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(200)
		// 	return
		// }

		c.Next()
	}
}
