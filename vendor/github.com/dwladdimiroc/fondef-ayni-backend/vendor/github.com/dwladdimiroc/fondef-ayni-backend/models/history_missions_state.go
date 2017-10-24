package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type History_missions_state struct {
	//gorm.Model
	Id    uint   `gorm:"primary_key"`
	State string `gorm:"column:state;not null;" json:"state"`
}

func HistoryMissionsStateCRUD(app *gin.Engine) {
	app.GET("/historyMissionsState/:id", HistoryMissionsStateFetchOne)
	app.GET("/historyMissionsState/", HistoryMissionsStateFetchAll)
	app.POST("/historyMissionsState/", HistoryMissionsStateCreate)
	app.PUT("/historyMissionsState/:id", HistoryMissionsStateUpdate)
	app.DELETE("/historyMissionsState/:id", HistoryMissionsStateRemove)
}

func HistoryMissionsStateFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var historyMissionsState History_missions_state
	if err := db.Find(&historyMissionsState, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, historyMissionsState)
	}
}

func HistoryMissionsStateFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var historyMissionsState []History_missions_state
	db.Find(&historyMissionsState)

	c.JSON(http.StatusOK, historyMissionsState)
}

func HistoryMissionsStateCreate(c *gin.Context) {
	var historyMissionsState History_missions_state
	e := c.BindJSON(&historyMissionsState)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&historyMissionsState).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, historyMissionsState)
	}
}

func HistoryMissionsStateUpdate(c *gin.Context) {
	var historyMissionsState History_missions_state
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&historyMissionsState).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&historyMissionsState)

		db.Save(&historyMissionsState)
		c.JSON(200, historyMissionsState)
	}
}

func HistoryMissionsStateRemove(c *gin.Context) {
	var historyMissionsState History_missions_state

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&historyMissionsState).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&historyMissionsState)
		c.JSON(http.StatusOK, historyMissionsState)
	}
}
