package controllers

import (
	"net/http"
	"path/filepath"
	"context"


	"github.com/Eco-Led/EcoLed-Back_test/services"
	"github.com/gin-gonic/gin"

)

type ProfileImageControllers struct{}


// UploadImage uploads image
func (ctr ProfileImageControllers) UploadProfileImage(c *gin.Context){
	var imageService services.ImageService
	
	//By form-data type, file is uploaded (in Controller)
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename := filepath.Base(file.Filename)
	//Open file (in Controller)
	filecontent, _ := file.Open()
	defer filecontent.Close()

	// Get userID from token & Chage type to uint (in Controller)
	userIDInterface, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get userIDInterface",
		})
		return
	}

	userIDInt64, ok := userIDInterface.(int64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to convert userID into int64",
		})
		return
	}
	userID := uint(userIDInt64)

	bucketName := "ecoled_test_profile_images"

	//Get imageURL (in Controller)
	imageURL, err := imageService.UploadImage(context.Background(), filecontent, userID, bucketName, filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Return imageURL (in Controller)
	c.JSON(http.StatusOK, gin.H{
		"Success to Upload!: %s": imageURL,
	})


}