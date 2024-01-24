package controllers

import (
	"net/http"
	"strconv"

	"github.com/Eco-Led/EcoLed-Back_test/services"

	"github.com/gin-gonic/gin"
)

type PaylogControllers struct{}

var paylogService = new(services.PaylogServices)

func (ctr PaylogControllers) CreatePaylog(c *gin.Context) {
	// Bind paylogForm from JSON
	var paylogForm services.PaylogForm
	if err := c.ShouldBindJSON(&paylogForm); err != nil {
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

	// Create paylog (service)
	err := paylogService.CreatePaylog(userID, paylogForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Success",
	})

}

func (ctr PaylogControllers) UpdatePaylog(c *gin.Context) {
	// Get paylogID from param
	paylogIDstring := c.Param("paylogID")
	paylogIDint64, err1 := strconv.ParseUint(paylogIDstring, 10, 64)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get paylogID",
		})
		return
	}
	paylogID := uint(paylogIDint64)


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

	// Get paylogForm from JSON
	var paylogForm services.PaylogForm
	if err2 := c.ShouldBindJSON(&paylogForm); err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err2.Error(),
		})
		return
	}

	// Update paylog (service)
	var err3 error
	err3 = paylogService.UpdatePaylog(userID, paylogID, paylogForm)
	if err3 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err3.Error(),
		})
		return
	}

	// Response paylog
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
	})

}

func (ctr PaylogControllers) DeletePaylog(c *gin.Context) {
	// Get paylogID from param
	paylogIDstring := c.Param("paylogID")
	paylogIDint64, err1 := strconv.ParseUint(paylogIDstring, 10, 64)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get paylogID",
		})
		return
	}
	paylogID := uint(paylogIDint64)

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

	// Delete paylog (service)
	err2 := paylogService.DeletePaylog(userID, paylogID)
	if err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err2.Error(),
		})
		return
	}

	// Response paylog
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})

}