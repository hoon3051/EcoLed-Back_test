package utils

import (
	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

// UploadImage uploads image to server

func UploadImage(c *gin.Context) {
	//By form-data type, file is uploaded
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//Get filename
	filename := filepath.Base(file.Filename)
	uniqueFilename := time.Now().Format("20060102150405") + "_" + filename

	//Save file to uploads directory(in server)
	if err := c.SaveUploadedFile(file, "uploads/"+uniqueFilename); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//Get imageURL
	imageURL := "http://localhost:8080/uploads/" + uniqueFilename

	//TODO : Save imageURL to DB in controller (return profile)
	profile := models.Profiles{Profile_image: imageURL}
	initializers.DB.Create(&profile)

	//TODO : in controller.
	c.JSON(http.StatusOK, gin.H{
		"파일 업로드 성공: %s": imageURL,
	})
}
