package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Volunteers_status struct {
	//gorm.Model
	Id     uint   `gorm:"primary_key"`
	Status string `gorm:"column:status;not null;" json:"status"`
}

func VolunteersStatusCRUD(app *gin.Engine) {
	app.GET("/volunteersStatus/:id", VolunteersStatusFetchOne)
	app.GET("/volunteersStatus/", VolunteersStatusFetchAll)
	app.POST("/volunteersStatus/", VolunteersStatusCreate)
	app.PUT("/volunteersStatus/:id", VolunteersStatusUpdate)
	app.DELETE("/volunteersStatus/:id", VolunteersStatusRemove)
}

func VolunteersStatusFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var volunteersStatus Volunteers_status
	if err := db.Find(&volunteersStatus, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, volunteersStatus)
	}
}

func VolunteersStatusFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var volunteersStatus []Volunteers_status
	db.Find(&volunteersStatus)

	c.JSON(http.StatusOK, volunteersStatus)
}

func VolunteersStatusCreate(c *gin.Context) {
	var volunteersStatus Volunteers_status
	e := c.BindJSON(&volunteersStatus)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&volunteersStatus).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, volunteersStatus)
	}
}

func VolunteersStatusUpdate(c *gin.Context) {
	var volunteersStatus Volunteers_status
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&volunteersStatus).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&volunteersStatus)

		db.Save(&volunteersStatus)
		c.JSON(200, volunteersStatus)
	}
}

func VolunteersStatusRemove(c *gin.Context) {
	var volunteersStatus Volunteers_status

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&volunteersStatus).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&volunteersStatus)
		c.JSON(http.StatusOK, volunteersStatus)
	}
}
