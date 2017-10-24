package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Volunteer struct {
	//gorm.Model
	Id                  uint      `gorm:"primary_key"`
	Volunteer_status_id int       `gorm:"column:volunteer_status_id;not null;" json:"volunteer_status_id"`
	Token               string    `gorm:"column:token;not null;" json:"token"`
	User_id             uint      `gorm:"column:user_id;" json:"user_id"`
	Abilities           []Ability `gorm:"many2many:volunteers_abilities;"`
}

func VolunteerCRUD(app *gin.Engine) {
	app.GET("/volunteer/:id", VolunteerFetchOne)
	app.GET("/volunteer/", VolunteerFetchAll)
	app.POST("/volunteer/", VolunteerCreate)
	app.PUT("/volunteer/:id", VolunteerUpdate)
	app.DELETE("/volunteer/:id", VolunteerRemove)
}

func VolunteerFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var volunteer Volunteer

	if err := db.Find(&volunteer, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Model(&volunteer).Related(&volunteer.Abilities, "Abilities")

		c.JSON(http.StatusOK, volunteer)
	}
}

func VolunteerFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var volunteers []Volunteer
	if err := db.Find(&volunteers).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		for i := range volunteers {
			db.Model(&volunteers[i]).Related(&volunteers[i].Abilities, "Abilities")
		}
		c.JSON(http.StatusOK, volunteers)
	}
}

func VolunteerCreate(c *gin.Context) {
	var volunteer Volunteer
	e := c.BindJSON(&volunteer)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&volunteer).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, volunteer)
	}
}

func VolunteerUpdate(c *gin.Context) {
	var volunteer Volunteer
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&volunteer).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&volunteer)

		db.Save(&volunteer)
		c.JSON(200, volunteer)
	}
}

func VolunteerRemove(c *gin.Context) {
	var volunteer Volunteer

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&volunteer).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&volunteer)
		c.JSON(http.StatusOK, volunteer)
	}
}
