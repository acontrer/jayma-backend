package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Emergencies_type struct {
	//gorm.Model
	Id   uint   `gorm:"primary_key"`
	Type string `gorm:"column:type;not null;" json:"type"`
}

func EmergenciesTypeCRUD(app *gin.Engine) {
	app.GET("/emType/:id", EmTypeFetchOne)
	app.GET("/emType/", EmTypeFetchAll)
	app.POST("/emType/", EmTypeCreate)
	app.PUT("/emType/:id", EmTypeUpdate)
	app.DELETE("/emType/:id", EmTypeRemove)
}

func EmTypeFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var emType Emergencies_type
	if err := db.Find(&emType, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, emType)
	}
}

func EmTypeFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var emTypes []Emergencies_type
	db.Find(&emTypes)

	c.JSON(http.StatusOK, emTypes)
}

func EmTypeCreate(c *gin.Context) {
	var emType Emergencies_type
	e := c.BindJSON(&emType)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&emType).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, emType)
	}
}

func EmTypeUpdate(c *gin.Context) {
	var emType Emergencies_type
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&emType).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&emType)

		db.Save(&emType)
		c.JSON(200, emType)
	}
}

func EmTypeRemove(c *gin.Context) {
	var emType Emergencies_type

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&emType).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&emType)
		c.JSON(http.StatusOK, emType)
	}
}
