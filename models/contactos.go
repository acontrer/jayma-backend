package models

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/dwladdimiroc/fondef-jayma-backend/db"
	"github.com/dwladdimiroc/fondef-jayma-backend/utils"
)

type Contactos struct {
	Usuarios_1_id int `gorm:"column:usuarios_1_id;not null;primary_key" json:"user_1_id"`
	Usuarios_2_id int `gorm:"column:usuarios_2_id;not null;primary_key" json:"user_2_id"`
}

func ContactoCRUD(app *gin.Engine) {
	app.GET("/contacto/:id", ContactoFetchOne)
	app.GET("/contacto/", ContactoFetchAll)
	app.POST("/contacto/", ContactoCreate)
	app.PUT("/contacto/:id", ContactoUpdate)
	app.DELETE("/contacto/:id", ContactoRemove)
}

func ContactoFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var contacto Contactos
	if err := db.Find(&contacto, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, contacto)
	}
}

func ContactoFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var contactos []Contactos
	if err := db.Find(&contactos).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, contactos)
	}
}

func ContactoCreate(c *gin.Context) {
	var contacto Contactos
	e := c.BindJSON(&contacto)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&contacto).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, contacto)
	}
}

func ContactoUpdate(c *gin.Context) {
	var contacto Contactos
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&contacto).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&contacto)

		db.Save(&contacto)
		c.JSON(http.StatusOK, contacto)
	}
}

func ContactoRemove(c *gin.Context) {
	var contacto Contactos

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&contacto).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&contacto)
		c.JSON(http.StatusOK, contacto)
	}
}