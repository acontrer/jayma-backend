package models

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/api"
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type History_missions struct {
	Mission_id               int `gorm:"column:mission_id;not null;primary_key;" json:"mission_id"`
	Mission                  Mission
	Volunteer_id             int `gorm:"column:volunteer_id;not null;primary_key;" json:"volunteer_id"`
	Volunteer                Volunteer
	History_mission_state_id int `gorm:"column:history_mission_state_id;not null;" json:"history_mission_state_id"`
	History_mission_state    History_missions_state
}

func HistoryMissionsCRUD(app *gin.Engine) {
	app.GET("/historyMissions/mission/:id", HistoryMissionsFetchMission)
	app.GET("/historyMissions/volunteer/:id", HistoryMissionsFetchVolunteer)
	app.GET("/historyMissions/", HistoryMissionsFetchAll)
	app.POST("/historyMissions/", HistoryMissionsCreate)
	app.PUT("/historyMissions/:mission/:volunteer", HistoryMissionsUpdate)
	app.DELETE("/historyMissions/:id", HistoryMissionsRemove)
}

func HistoryMissionsFetchMission(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var historyMission []History_missions
	db.Where("mission_id = ?", id).Find(&historyMission)

	type Response struct {
		Users User `json:"Users"`
		State int  `json:"state"`
	}

	var response []Response
	response = make([]Response, len(historyMission))

	for i := range historyMission {
		db.Model(&historyMission[i]).Related(&historyMission[i].Volunteer, "Volunteer_id")
		var user User
		db.Model(&historyMission[i].Volunteer).Related(&user, "Volunteer_id")

		response[i].Users = user
		response[i].State = historyMission[i].History_mission_state_id
	}

	c.JSON(http.StatusOK, response)
}

func HistoryMissionsFetchVolunteer(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var historyMission []History_missions
	db.Where("volunteer_id = ?", id).Find(&historyMission)

	type Response struct {
		Missions []Mission `json:"Missions"`
	}

	var response Response
	response.Missions = make([]Mission, len(historyMission))

	for i := range historyMission {
		var mission Mission
		db.Model(&historyMission[i]).Related(&mission, "Mission_id")
		response.Missions[i] = mission
	}

	c.JSON(http.StatusOK, response)
}

func HistoryMissionsFetchAll(c *gin.Context) {
	db := db.Database()
	defer db.Close()

	var historyMissions []History_missions
	db.Find(&historyMissions)

	c.JSON(http.StatusOK, historyMissions)
}

func HistoryMissionsCreate(c *gin.Context) {
	var historyMissions History_missions
	e := c.BindJSON(&historyMissions)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&historyMissions).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		db.Model(&historyMissions).Related(&historyMissions.Volunteer, "Volunteer_id")
		db.Model(&historyMissions).Related(&historyMissions.Mission, "Mission_id")

		titulo := "Te han invitado a una nueva misión"
		notificacion := "Te han invitado a la misión " + historyMissions.Mission.Title

		api.SendNotification(historyMissions.Volunteer.Token, titulo, notificacion)
		c.JSON(http.StatusCreated, historyMissions)
	}
}

func HistoryMissionsUpdate(c *gin.Context) {
	var historyMissions History_missions
	missionId := c.Param("mission")
	volunteerId := c.Param("volunteer")

	db := db.Database()
	defer db.Close()

	if err := db.Where("mission_id = ? AND volunteer_id = ?", missionId, volunteerId).First(&historyMissions).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&historyMissions)

		db.Save(&historyMissions)
		c.JSON(http.StatusOK, historyMissions)
	}
}

func HistoryMissionsRemove(c *gin.Context) {
	var historyMissions History_missions
	missionId := c.Param("mission")
	volunteerId := c.Param("volunteer")

	db := db.Database()
	defer db.Close()

	if err := db.Where("mission_id = ? AND volunteer_id = ?", missionId, volunteerId).First(&historyMissions).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&historyMissions)

		db.Delete(&historyMissions)
		c.JSON(http.StatusOK, historyMissions)
	}
}
