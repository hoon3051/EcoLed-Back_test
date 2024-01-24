package services

import (
	"context"
	"errors"
	"io"
	"os"
	"fmt"
	"time"
	

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"

)

type ImageService struct{}

func (srv ImageService) UploadImage(ctx context.Context, file io.Reader, userID uint, bucketName string, fileName string) (imageURL string, err error) {
	//Get filename (in Service)
	uniqueFilename := time.Now().Format("20060102150405") + "_" + fileName
	
	// Upload image on Google Cloud Storage
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS1")))
	if err != nil {
		err = errors.New("Failed to create client")
		return "", err
	}
	defer client.Close()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucketName).Object(uniqueFilename).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		err = errors.New("Failed to upload image1")
		return "", err
	}

	if err := wc.Close(); err != nil {
		err = errors.New("Failed to upload image2")
		return "", err
	}

	//Get imageURL
	imageURL = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, uniqueFilename)

	//Save imageURL to DB (BY BUCKET NAME)
	switch bucketName {
		case "ecoled_test_profile_images": 
			result := initializers.DB.Model(&models.Profiles{}).
			Where("user_id = ?", userID).
			Update("profile_image", imageURL)

			// 에러 처리
			if result.Error != nil {
				err = errors.New("Failed to upload image in DB")
				return imageURL, err
			}
							
			
		case "ecoled_test_post_images": 
			var post models.Posts
			result1 := initializers.DB.Where("user_id = ?", userID).Order("updated_at DESC").First(&post)
			if result1.Error != nil {
				err = errors.New("Failed to get last created post")
				return imageURL, err
			}

			post.Image = imageURL
			result2 := initializers.DB.Save(&post)
			if result2.Error != nil {
				err = errors.New("Failed to upload image in DB")
				return imageURL, err
			}

	}

	return imageURL, nil

}