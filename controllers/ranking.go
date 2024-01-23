package controllers

import (
	"github.com/Eco-Led/EcoLed-Back_test/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RankingControllers struct{}

var rankingService = new(services.RankingServices)

func (ctr RankingControllers) GetRanking(c *gin.Context) {
	// Get ranking (service)
	ranking, err := rankingService.GetRanking()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ranking": ranking,
	})
}