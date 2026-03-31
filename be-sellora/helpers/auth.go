package helpers

import (
	"be-sellora/database"
	"be-sellora/models"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetAuthUserID(c *gin.Context) (uint, error) {
	username, exists := c.Get("username")
	if !exists {
		return 0, errors.New("user not found in context")
	}

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}
