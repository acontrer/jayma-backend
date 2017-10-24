package models

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"time"

	"github.com/dwladdimiroc/fondef-ayni-backend/api"
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Mission struct {
	//gorm.Model
	Id                      uint      `gorm:"primary_key"`
	CreatedAt               time.Time `gorm:"column:createAt" json:"createAt"`
	Meeting_point_latitude  float64   `gorm:"column:meeting_point_latitude;not null;" json:"meeting_point_latitude"`
	Meeting_point_longitude float64   `gorm:"column:meeting_point_longitude;not null;" json:"meeting_point_longitude"`
	Title                   string    `gorm:"column:title;not null;" json:"title"`
	Description             string    `gorm:"column:description;not null;" json:"description"`
	Meeting_point_address   string    `gorm:"column:meeting_point_address;not null;" json:"meeting_point_address"`
	Start_date              time.Time `gorm:"column:start_date;not null;" json:"start_date"`
	Finish_date             time.Time `gorm:"column:finish_date;not null;" json:"finish_date"`
	Scheduled_start_date    time.Time `gorm:"column:scheduled_start_date;not null;" json:"scheduled_start_date"`
	Scheduled_finish_date   time.Time `gorm:"column:scheduled_finish_date;not null;" json:"scheduled_finish_date"`
	Assertiveness_text      float64   `gorm:"column:assertiveness_text;not null;" json:"assertiveness_text"`
	Emergency_id            int       `gorm:"column:emergency_id;not null;" json:"emergency_id"`
	Emergency               Emergency `gorm:"ForeignKey:Emergency_id;AssociationForeignKey:Id"`
	Mission_status_id       int       `gorm:"column:mission_status_id;not null;" json:"mission_status_id"`
	User_id                 int       `gorm:"column:user_id;not null;" json:"user_id"`
	User                    User      `gorm:"ForeignKey:User_id;AssociationForeignKey:Id"`
	Abilities               []Ability `gorm:"many2many:missions_abilities;"`
	Problems                []Problem
	Files                   []File
}

func MissionCRUD(app *gin.Engine) {
	app.GET("/mission/:id", MissionFetchOne)
	app.GET("/mission/", MissionFetchAll)
	app.POST("/mission/", MissionCreate)
	app.PUT("/mission/:id", MissionUpdate)
	app.DELETE("/mission/:id", MissionRemove)
}

func MissionFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var mission Mission
	if err := db.Find(&mission, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Model(&mission).Related(&mission.User, "Users_id")
		db.Model(&mission.User).Related(&mission.User.Volunteer, "Users_id")
		db.Model(&mission).Related(&mission.Abilities, "Abilities")
		db.Model(&mission).Related(&mission.Problems, "Missions_id")
		db.Model(&mission).Related(&mission.Files, "Missions_id")
		db.Model(&mission).Related(&mission.Emergency, "Emergencies_id")

		for i := range mission.Problems {
			db.Model(&mission.Problems[i]).Related(&mission.Problems[i].Answer, "Problem_id")
		}

		c.JSON(http.StatusOK, mission)
	}
}

func MissionFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var missions []Mission
	db.Find(&missions)

	for i := range missions {
		db.Model(&missions[i]).Related(&missions[i].User, "Users_id")
		db.Model(&missions[i]).Related(&missions[i].Abilities, "Abilities")
		db.Model(&missions[i]).Related(&missions[i].Problems, "Missions_id")
		db.Model(&missions[i]).Related(&missions[i].Emergency, "Emergencies_id")
	}

	c.JSON(http.StatusOK, missions)
}

func MissionCreate(c *gin.Context) {
	var mission Mission
	if err := c.BindJSON(&mission); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		address, err := api.MissionReverseGeocodingPoint(mission.Meeting_point_latitude, mission.Meeting_point_longitude)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			mission.Meeting_point_address = address
			mission.Assertiveness_text = utils.Asertiveness(mission.Description)

			db := db.Database()
			defer db.Close()

			if err := db.Create(&mission).Error; err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			} else {
				c.JSON(http.StatusCreated, mission)
			}
		}
	}
}

func MissionUpdate(c *gin.Context) {
	var mission Mission
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&mission).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&mission)

		db.Save(&mission)
		c.JSON(200, mission)
	}
}

func MissionRemove(c *gin.Context) {
	var mission Mission

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&mission).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&mission)
		c.JSON(http.StatusOK, mission)
	}
}
