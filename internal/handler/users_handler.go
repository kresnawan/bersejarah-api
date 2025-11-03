package handler

import (
	"net/http"

	"app/internal/storage"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	// data := storage.DbConnect()
	// c.Data(200, "application/json", data)
}

func AddUser(c *gin.Context) {
	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	var ReqBody RequestBody

	if err := c.BindJSON(&ReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Body accepted",
		"username": ReqBody.Username,
		"password": ReqBody.Password,
		"name":     ReqBody.Name,
	})

	insertId, err := storage.RegisterUser(ReqBody.Name, ReqBody.Username, ReqBody.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed insert user data"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered, id: " + string(insertId)})
}

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}
