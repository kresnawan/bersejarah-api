package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"app/internal/model"
	"app/internal/storage"
	"app/utility"

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
		AlamatJalan     string  `json:"alamat_jalan"`
		AlamatDusun     string  `json:"alamat_dusun"`
		AlamatKelurahan string  `json:"alamat_kelurahan"`
		AlamatKecamatan string  `json:"alamat_kecamatan"`
		AlamatKabupaten string  `json:"alamat_kabupaten"`
		DeskripsiTempat string  `json:"deskripsi"`
		Latitude        float32 `json:"lat"`
		Longitude       float32 `json:"long"`
	}

	var ReqBody RequestBody

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed reading request body"})
		c.Abort()
		return
	}

	var address string = utility.ParseAddress(
		ReqBody.AlamatJalan,
		ReqBody.AlamatDusun,
		ReqBody.AlamatKelurahan,
		ReqBody.AlamatKecamatan,
		ReqBody.AlamatKabupaten,
	)

	insertId, err := storage.InsertDataTempat(
		ReqBody.NamaTempat,
		address,
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
	var imageArr []model.FileD
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed parse data id for image", "err": err})
	}

	// Memeriksa apakah ada file yang ukurannya lebih dari 1MB

	// for _, file := range files {
	// 	if file.Size > 5000000 {
	// 		c.Data(http.StatusBadRequest, "application/json", []byte(`{"message":"Images size too big"}`))
	// 		c.Abort()
	// 		return
	// 	}
	// }

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

		var fileIns model.FileD
		fileIns.Filename = filename
		fileIns.Size = int(file.Size)

		imageArr = append(imageArr, fileIns)
		counter++
	}

	result, err := storage.UploadFotoTempat(id, imageArr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed insert images into database", "err": err})
		c.Abort()
		return
	}

	type SuccessResponse struct {
		Message string        `json:"message"`
		Files   []model.FileD `json:"files"`
		Query   string        `json:"query"`
	}

	var response SuccessResponse
	response.Message = fmt.Sprintf("Successfully saved %v files", counter)
	response.Files = imageArr
	response.Query = string(result)

	resoo, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed convert image array to json", "err": err})
	}

	c.Data(200, "application/json", resoo)
	c.Abort()
	return

}

func DeleteDataTempat(c *gin.Context) {

}

func UpdateDataTempat(c *gin.Context) {

}
