package controllers

import (
	"net/http"

	"github.com/Eco-Led/EcoLed-Back_test/services"

	"github.com/gin-gonic/gin"
)

type ProfileControllers struct{}

var profileService = new(services.ProfileServices)

func (ctr ProfileControllers) UpdateProfile(c *gin.Context) {
	// Bind JSON
	var profileForm services.ProfileForm
	if err := c.ShouldBindJSON(&profileForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
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

	// Update (service)
	err := profileService.UpdateProfile(userID, profileForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
	})

}

func (ctr ProfileControllers) GetProfile(c *gin.Context) {
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

	// Get profile (service)
	profile, err := profileService.GetProfile(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
	
}