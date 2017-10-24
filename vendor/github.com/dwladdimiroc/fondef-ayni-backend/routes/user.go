package routes

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

func FetchUserInformacion(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		var user models.User
		if err := user.UserFetchEmail(userId.(string)); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.String(http.StatusInternalServerError, "Error parser Token")
	}
}

func EditUserInformacion(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		var user models.User

		db := db.Database()
		defer db.Close()

		if err := user.UserFetchEmail(userId.(string)); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.BindJSON(&user)

			var copyUser models.User
			copyUser = user
			db.Model(&user.Volunteer).Association("Abilities").Clear()

			db.Save(&copyUser)

			c.JSON(http.StatusOK, copyUser)
		}
	} else {
		c.String(http.StatusInternalServerError, "Error parser Token")
	}
}

func EditTokenUser(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		var user models.User

		db := db.Database()
		defer db.Close()

		if err := user.UserFetchEmail(userId.(string)); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.BindJSON(&user.Volunteer)

			db.Save(&user.Volunteer)

			c.JSON(http.StatusOK, user)
		}
	} else {
		c.String(http.StatusInternalServerError, "Error parser Token")
	}
}
