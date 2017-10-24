package routes

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

func VoluntariosDisponibles(c *gin.Context) {
	db := db.Database()
	defer db.Close()

	var volunteers []models.Volunteer
	db.Where("volunteer_status_id = ?", 1).Find(&volunteers)

	users := make([]models.User, len(volunteers))
	for i := range volunteers {
		db.Model(&volunteers[i]).Related(&users[i], "User_id")
	}

	c.JSON(http.StatusOK, users)
}

func VoluntarioInteres(c *gin.Context) {
	var historyMissions models.History_missions
	e := c.BindJSON(&historyMissions)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&historyMissions).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	var volunteer models.Volunteer
	if err := db.Where("id = ?", historyMissions.Volunteer_id).First(&volunteer).Error; err != nil {
		if err := db.Where("volunteer_id = ?", historyMissions.Volunteer_id).First(&historyMissions).Error; err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {
			db.Delete(&historyMissions)
		}

		c.String(http.StatusNotFound, err.Error())
	} else {
		volunteer.Volunteer_status_id = 2
		db.Save(&volunteer)
	}

	type Response struct {
		HistoryMissions models.History_missions `json:"HistoryMissions"`
		Volunteer       models.Volunteer        `json:"Volunteer"`
	}

	var resp Response
	resp.HistoryMissions = historyMissions
	resp.Volunteer = volunteer

	c.JSON(http.StatusOK, resp)
}
