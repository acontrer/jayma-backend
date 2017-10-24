package routes

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
	//	"github.com/dwladdimiroc/fondef-ayni-backend/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

func InvitacionesPorUsuario(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		if email, ok := userId.(string); ok {
			var user models.User
			if err := user.UserFetchEmail(email); err != nil {
				c.String(http.StatusNotFound, err.Error())
			} else {
				db := db.Database()
				defer db.Close()

				db.Model(&user).Related(&user.Volunteer, "Users_id")

				var histories []models.History_missions
				if err := db.Where("volunteer_id = ? AND history_mission_state_id = 1", user.Volunteer.Id).Find(&histories).Error; err != nil {
					c.String(http.StatusNotFound, err.Error())
				} else {
					for i := range histories {
						db.Model(&histories[i]).Related(&histories[i].Mission, "Mission_id")
					}

					c.JSON(http.StatusOK, histories)
				}
			}
		} else {
			c.String(http.StatusInternalServerError, "Error parse string email")
		}
	} else {
		c.String(http.StatusInternalServerError, "Error parser Token")
	}
}

func MisionActivaUsuario(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		if email, ok := userId.(string); ok {
			var user models.User
			if err := user.UserFetchEmail(email); err != nil {
				c.String(http.StatusNotFound, err.Error())
			} else {
				db := db.Database()
				defer db.Close()

				db.Model(&user).Related(&user.Volunteer, "Users_id")

				var histories []models.History_missions
				if err := db.Where("volunteer_id = ? AND history_mission_state_id = 2 OR history_mission_state_id = 3", user.Volunteer.Id).Find(&histories).Error; err != nil {
					c.String(http.StatusNotFound, err.Error())
				} else {
					for i := range histories {
						db.Model(&histories[i]).Related(&histories[i].Mission, "Mission_id")
						db.Model(&histories[i].Mission).Related(&histories[i].Mission.User, "Users_id")
						db.Model(&histories[i].Mission.User).Related(&histories[i].Mission.User.Volunteer, "Users_id")
						db.Model(&histories[i].Mission).Related(&histories[i].Mission.Abilities, "Abilities")
						db.Model(&histories[i].Mission).Related(&histories[i].Mission.Problems, "Missions_id")
						db.Model(&histories[i].Mission).Related(&histories[i].Mission.Files, "Missions_id")
						db.Model(&histories[i].Mission).Related(&histories[i].Mission.Emergency, "Emergencies_id")
					}

					c.JSON(http.StatusOK, histories)
				}
			}
		} else {
			c.String(http.StatusInternalServerError, "Error parse string email")
		}
	} else {
		c.String(http.StatusInternalServerError, "Error parser Token")
	}
}

func ObtenerEstadoUsuario(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		if email, ok := userId.(string); ok {
			var user models.User
			if err := user.UserFetchEmail(email); err != nil {
				c.String(http.StatusNotFound, err.Error())
			} else {
				missionId := c.Param("mission")
				var history models.History_missions

				db := db.Database()
				defer db.Close()

				db.Model(&user).Related(&user.Volunteer, "Users_id")

				if err := db.Where("mission_id = ? AND volunteer_id = ?", missionId, user.Volunteer.Id).First(&history).Error; err != nil {
					c.String(http.StatusNotFound, err.Error())
				} else {
					c.JSON(http.StatusOK, history)
				}
			}

		} else {
			c.String(http.StatusInternalServerError, "Error parse string email")
		}
	} else {
		c.String(http.StatusInternalServerError, "Error parser Token")
	}
}

func CambiarEstadoUsuario(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		if email, ok := userId.(string); ok {
			var user models.User
			if err := user.UserFetchEmail(email); err != nil {
				c.String(http.StatusNotFound, err.Error())
			} else {
				missionId := c.Param("mission")
				var history models.History_missions

				db := db.Database()
				defer db.Close()

				db.Model(&user).Related(&user.Volunteer, "Users_id")

				if err := db.Where("mission_id = ? AND volunteer_id = ?", missionId, user.Volunteer.Id).First(&history).Error; err != nil {
					c.String(http.StatusNotFound, err.Error())
				} else {
					c.BindJSON(&history)

					db.Save(&history)
					c.JSON(http.StatusOK, history)
				}
			}

		} else {
			c.String(http.StatusInternalServerError, "Error parse string email")
		}
	} else {
		c.String(http.StatusInternalServerError, "Error parser Token")
	}
}
