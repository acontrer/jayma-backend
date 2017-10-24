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

type File struct {
	//	gorm.Model
	Id         uint   `gorm:"primary_key"`
	Mission_id int    `gorm:"column:mission_id;not null;" json:"mission_id"`
	File       string `gorm:"column:file;not null;" json:"file"`
}

func FileCRUD(app *gin.Engine) {
	app.GET("/file/:id", FileFetchOne)
	app.GET("/file/", FileFetchAll)
	app.POST("/file/", FileCreate)
	app.PUT("/file/:id", FileUpdate)
	app.DELETE("/file/:id", FileRemove)
}

func FileFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var file File
	if err := db.Find(&file, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, file)
	}
}

func FileFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var files []File
	if err := db.Find(&files).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {

		c.JSON(http.StatusOK, files)
	}
}

func FileCreate(c *gin.Context) {
	var file File
	e := c.BindJSON(&file)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&file).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, file)
	}
}

func FileUpdate(c *gin.Context) {
	var file File
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&file).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&file)

		db.Save(&file)
		c.JSON(200, file)
	}
}

func FileRemove(c *gin.Context) {
	var file File

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&file).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&file)
		c.JSON(http.StatusOK, file)
	}
}
