package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Missions_status struct {
	//gorm.Model
	Id     uint   `gorm:"primary_key"`
	Status string `gorm:"column:status;not null;" json:"status"`
}

func MissionsStatusCRUD(app *gin.Engine) {
	app.GET("/missionsStatus/:id", MissionsStatusFetchOne)
	app.GET("/missionsStatus/", MissionsStatusFetchAll)
	app.POST("/missionsStatus/", MissionsStatusCreate)
	app.PUT("/missionsStatus/:id", MissionsStatusUpdate)
	app.DELETE("/missionsStatus/:id", MissionsStatusRemove)
}

func MissionsStatusFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var missionsStatus Missions_status
	if err := db.Find(&missionsStatus, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, missionsStatus)
	}
}

func MissionsStatusFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var missionsStatus []Missions_status
	db.Find(&missionsStatus)

	c.JSON(http.StatusOK, missionsStatus)
}

func MissionsStatusCreate(c *gin.Context) {
	var missionsStatus Missions_status
	e := c.BindJSON(&missionsStatus)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&missionsStatus).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, missionsStatus)
	}
}

func MissionsStatusUpdate(c *gin.Context) {
	var missionsStatus Missions_status
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&missionsStatus).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&missionsStatus)

		db.Save(&missionsStatus)
		c.JSON(200, missionsStatus)
	}
}

func MissionsStatusRemove(c *gin.Context) {
	var missionsStatus Missions_status

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&missionsStatus).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&missionsStatus)
		c.JSON(http.StatusOK, missionsStatus)
	}
}
