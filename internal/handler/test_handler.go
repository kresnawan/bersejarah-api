package handler

import (
	"app/utility"

	"github.com/gin-gonic/gin"
)

func GetHashFromParams(c *gin.Context) {
	stringToHash := c.Param("text")

	result, err := utility.HashPasswordSecure(stringToHash)
	if err != nil {
		c.JSON(400, gin.H{"message": "Error hashing"})
	}

	c.JSON(200, gin.H{"message": stringToHash, "hash": result})

}
