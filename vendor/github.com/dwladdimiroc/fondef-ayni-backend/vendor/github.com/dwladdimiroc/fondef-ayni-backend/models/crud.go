package models

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	AbilityCRUD(app)
	AnswerCRUD(app)
	EmergenciesStatusCRUD(app)
	EmergenciesTypeCRUD(app)
	EmergencyCRUD(app)
	FileCRUD(app)
	HistoryMissionsCRUD(app)
	HistoryMissionsStateCRUD(app)
	MissionCRUD(app)
	MissionsStatusCRUD(app)
	ProblemCRUD(app)
	UserCRUD(app)
	UserTypeCRUD(app)
	VolunteerCRUD(app)
	VolunteersStatusCRUD(app)
}
