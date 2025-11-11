package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
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
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed reading the image"})
		c.Abort()
		return
	}

	files := form.File["images"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No files uploaded"})
		c.Abort()
		return
	}

	var counter = 0

	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		now := time.Now()

		epochMill := now.UnixNano()
		filename := strconv.FormatInt(epochMill, 10) + fileExt
		dst := filepath.Join("./uploads", filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed saving the files in the cloud", "err": err})
			c.Abort()
			return
		}

		counter++
	}

	c.JSON(200, gin.H{"message": fmt.Sprintf("Successfully saved %v files", counter)})

}

func DeleteDataTempat(c *gin.Context) {

}

func UpdateDataTempat(c *gin.Context) {

}
