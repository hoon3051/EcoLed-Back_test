package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
	"context"
	"io"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
	"github.com/gin-gonic/gin"
)

type ImageControllers struct{}

// UploadImage uploads image
func (ctr ImageControllers) UploadImage(c *gin.Context){
	//By form-data type, file is uploaded
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}



	//Get filename
	filename := filepath.Base(file.Filename)
	uniqueFilename := time.Now().Format("20060102150405") + "_" + filename

	//Setting Google Cloud Storage 
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS1")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer client.Close()

	// Upload image on Google Cloud Storage
	bucket := "ecoled_test_profile_images"
	wc := client.Bucket(bucket).Object(uniqueFilename).NewWriter(ctx)
	filecontent, _ := file.Open()
	defer filecontent.Close()

	if _, err := io.Copy(wc, filecontent); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := wc.Close(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}


	// Get userID from token & Chage type to uint
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

	//Get imageURL
	imageURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, uniqueFilename)

	//TODO : Save imageURL to DB in controller (return profile)
	result := initializers.DB.Model(&models.Profiles{}).Where("user_id = ?", userID).Update("profile_image", imageURL)

    // 에러 처리
    if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
    }
   

	//TODO : in controller.
	c.JSON(http.StatusOK, gin.H{
		"Success to Upload!: %s": imageURL,
	})


}