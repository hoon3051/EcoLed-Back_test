package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Eco-Led/EcoLed-Back_test/services"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		tokenString := c.GetHeader("Authorization")

		// Check token
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized1",
			})
			return
		}

		// Get userID from token
		userID, err := services.TokenServices{}.ExtractTokenID(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Set userID to context
		c.Set("user_id", userID)

		c.Next()
	}
}