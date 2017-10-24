package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
)

func AddPermission(users ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("authGroups", users)
	}
}

func typeUser(userId string) int {
	db := db.Database()
	defer db.Close()

	var user models.User

	if err := db.Where("email = ?", userId).First(&user).Error; err != nil {
		return 0
	} else {
		return user.User_type_id
	}
}

func authorizatorUser(typeUser int, authGroups []int) bool {
	for _, authGroup := range authGroups {
		if authGroup == typeUser {
			return true
		}
	}
	return false
}
