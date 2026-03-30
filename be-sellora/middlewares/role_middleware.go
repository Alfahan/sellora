package middlewares

import (
	"be-sellora/database"
	"be-sellora/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Role(roleName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, exists := c.Get("username")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		var user models.User

		if err := database.DB.Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}

		hasRole := false
		for _, role := range user.Roles {
			if role.Name == roleName {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": http.StatusText(http.StatusForbidden)})
			c.Abort()
			return
		}

		c.Next()
	}
}
