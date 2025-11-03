package handler

import (
	"net/http"

	"app/internal/storage"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var ReqBody RequestBody

	// Membaca request body
	if err := c.BindJSON(&ReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Gagal membaca request body", "error": err.Error()})
		return
	}

	// DB error
	arr, err := storage.GetUserWithSpecificUsername(ReqBody.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB query failed", "error": err.Error()})
		c.Abort()
		return
	}

	// Username not found / array length = 0
	if len(arr) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Username not found"})
		c.Abort()
		return
	}

	// Check password
	match, err := argon2id.ComparePasswordAndHash(ReqBody.Password, arr[0].Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password compare error"})
		c.Abort()
		return
	}

	// Wrong password
	if match == false {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil"})
	c.Abort()
	return
}

func Register(c *gin.Context) {
	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	var ReqBody RequestBody

	// Membaca request body
	if err := c.BindJSON(&ReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Gagal membaca request body"})
		return
	}

	// Hashing password
	hashedPassword, err := argon2id.CreateHash(ReqBody.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Hashing password error", "error": err.Error()})
		c.Abort()
		return
	}

	// Insert data
	insertId, err := storage.RegisterUser(ReqBody.Name, ReqBody.Username, hashedPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username telah dipakai"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered, id: " + string(insertId)})
	c.Abort()
	return
}
