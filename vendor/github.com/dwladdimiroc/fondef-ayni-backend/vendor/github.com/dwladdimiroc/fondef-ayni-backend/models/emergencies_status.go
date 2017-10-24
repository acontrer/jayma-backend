package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Emergencies_status struct {
	//gorm.Model
	Id     uint   `gorm:"primary_key"`
	Status string `gorm:"column:status;not null;" json:"status"`
}

func EmergenciesStatusCRUD(app *gin.Engine) {
	app.GET("/emStatus/:id", EmStatusFetchOne)
	app.GET("/emStatus/", EmStatusFetchAll)
	app.POST("/emStatus/", EmStatusCreate)
	app.PUT("/emStatus/:id", EmStatusUpdate)
	app.DELETE("/emStatus/:id", EmStatusRemove)
}

func EmStatusFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var emStatus Emergencies_status
	if err := db.Find(&emStatus, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, emStatus)
	}
}

func EmStatusFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var emStatus []Emergencies_status
	db.Find(&emStatus)

	c.JSON(http.StatusOK, emStatus)
}

func EmStatusCreate(c *gin.Context) {
	var emStatus Emergencies_status
	e := c.BindJSON(&emStatus)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&emStatus).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, emStatus)
	}
}

func EmStatusUpdate(c *gin.Context) {
	var emStatus Emergencies_status
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&emStatus).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&emStatus)

		db.Save(&emStatus)
		c.JSON(200, emStatus)
	}
}

func EmStatusRemove(c *gin.Context) {
	var emStatus Emergencies_status

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&emStatus).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&emStatus)
		c.JSON(http.StatusOK, emStatus)
	}
}
