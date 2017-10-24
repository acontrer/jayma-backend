package auth

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func StatusPermission(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
	})
}

func AddPermission(users ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("authGroups", users)
	}
}

func typeUser(userId string) int {
	return 2
}

func authorizatorUser(typeUser int, authGroups []int) bool {
	for _, authGroup := range authGroups {
		if authGroup == typeUser {
			return true
		}
	}
	return false
}
