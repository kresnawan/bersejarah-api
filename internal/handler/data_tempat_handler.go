package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"app/internal/storage"

	"github.com/gin-gonic/gin"
)

func GetAllDataTempat(c *gin.Context) {
	data, err := storage.GetAllDataTempat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed fetch the data", "err": err})
		c.Abort()
		return
	}

	c.Data(200, "application/json", data)
}

func AddDataTempat(c *gin.Context) {
	type RequestBody struct {
		NamaTempat      string  `json:"nama_tempat"`
		AlamatTempat    string  `json:"alamat_tempat"`
		DeskripsiTempat string  `json:"deskripsi"`
		Latitude        float32 `json:"lat"`
		Longitude       float32 `json:"long"`
		Foto            string  `json:"foto"`
	}

	var ReqBody RequestBody

	if err := c.BindJSON(&ReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed reading request body"})
		c.Abort()
		return
	}

	insertId, err := storage.InsertDataTempat(
		ReqBody.NamaTempat,
		ReqBody.AlamatTempat,
		ReqBody.DeskripsiTempat,
		ReqBody.Latitude,
		ReqBody.Longitude,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed inserting data", "err": err})
		c.Abort()
		return
	}

	stringedInsertId := strconv.FormatInt(insertId, 10)

	c.JSON(200, gin.H{"message": "Data successfully inserted, insert ID: " + stringedInsertId})
}

func UploadFoto(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed reading the image"})
		c.Abort()
		return
	}

	defer file.Close()

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), fileExt)
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt

	dst := filepath.Join("./uploads", filename)
	if err := c.SaveUploadedFile(header, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed saving the file"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Successfully accept and save the image"})

}

func DeleteDataTempat(c *gin.Context) {

}

func UpdateDataTempat(c *gin.Context) {

}
