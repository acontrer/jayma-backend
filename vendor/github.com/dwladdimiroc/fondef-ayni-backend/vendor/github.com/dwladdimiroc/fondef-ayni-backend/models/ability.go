package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
	//	"fmt"
)

type Ability struct {
	//	gorm.Model
	Id      uint   `gorm:"primary_key"`
	Ability string `gorm:"column:ability;not null;" json:"ability"`
}

func AbilityCRUD(app *gin.Engine) {
	app.GET("/ability/:id", AbilityFetchOne)
	app.GET("/ability/", AbilityFetchAll)
	app.POST("/ability/", AbilityCreate)
	app.PUT("/ability/:id", AbilityUpdate)
	app.DELETE("/ability/:id", AbilityRemove)
}

func AbilityFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var ability Ability
	if err := db.Find(&ability, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, ability)
	}
}

func AbilityFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var abilities []Ability
	if err := db.Find(&abilities).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, abilities)
	}
}

func AbilityCreate(c *gin.Context) {
	var ability Ability
	e := c.BindJSON(&ability)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&ability).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, ability)
	}
}

func AbilityUpdate(c *gin.Context) {
	var ability Ability
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&ability).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&ability)

		db.Save(&ability)
		c.JSON(http.StatusOK, ability)
	}
}

func AbilityRemove(c *gin.Context) {
	var ability Ability

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&ability).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&ability)
		c.JSON(http.StatusOK, ability)
	}
}
