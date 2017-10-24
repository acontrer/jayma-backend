package models

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/dwladdimiroc/fondef-jayma-backend/db"
	"github.com/dwladdimiroc/fondef-jayma-backend/utils"
)

type Estados struct {
	Id     int    `gorm:"column:id;not null;" json:"id"`
	Estado string `gorm:"column:estado;not null;" json:"estado"`
}

func EstadoCRUD(app *gin.Engine) {
	app.GET("/estado/:id", EstadoFetchOne)
	app.GET("/estado/", EstadoFetchAll)
	app.POST("/estado/", EstadoCreate)
	app.PUT("/estado/:id", EstadoUpdate)
	app.DELETE("/estado/:id", EstadoRemove)
}

func EstadoFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var estado Estados
	if err := db.Find(&estado, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, estado)
	}
}

func EstadoFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var estados []Estados
	if err := db.Find(&estados).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, estados)
	}
}

func EstadoCreate(c *gin.Context) {
	var estado Estados
	e := c.BindJSON(&estado)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&estado).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, estado)
	}
}

func EstadoUpdate(c *gin.Context) {
	var estado Estados
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&estado).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&estado)

		db.Save(&estado)
		c.JSON(http.StatusOK, estado)
	}
}

func EstadoRemove(c *gin.Context) {
	var estado Estados

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&estado).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&estado)
		c.JSON(http.StatusOK, estado)
	}
}