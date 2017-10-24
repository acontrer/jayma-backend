package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Users_type struct {
	//gorm.Model
	Id   uint   `gorm:"primary_key"`
	Type string `gorm:"column:type;not null;" json:"type"`
}

func UserTypeCRUD(app *gin.Engine) {
	app.GET("/userType/:id", UserTypeFetchOne)
	app.GET("/userType/", UserTypeFetchAll)
	app.POST("/userType/", UserTypeCreate)
	app.PUT("/userType/:id", UserTypeUpdate)
	app.DELETE("/userType/:id", userTypeRemove)
}

func UserTypeFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var userType Users_type
	if err := db.Find(&userType, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, userType)
	}
}

func UserTypeFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var userTypes []Users_type
	db.Find(&userTypes)

	c.JSON(http.StatusOK, userTypes)
}

func UserTypeCreate(c *gin.Context) {
	var userType Users_type
	e := c.BindJSON(&userType)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&userType).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, userType)
	}
}

func UserTypeUpdate(c *gin.Context) {
	var userType Users_type
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&userType).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&userType)

		db.Save(&userType)
		c.JSON(200, userType)
	}
}

func userTypeRemove(c *gin.Context) {
	var userType Users_type

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&userType).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&userType)
		c.JSON(http.StatusOK, userType)
	}
}
