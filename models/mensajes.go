package models

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"time"

	"github.com/dwladdimiroc/fondef-jayma-backend/db"
	"github.com/dwladdimiroc/fondef-jayma-backend/utils"
)

type Mensajes struct {
	Id          int64     `gorm:"column:id;not null;" json:"id"`
	Mensaje     string    `gorm:"column:mensaje;not null;" json:"mensaje"`
	CreatedAt    time.Time `gorm:"column:fecha_creacion;not null;" json:"fecha_creacion"`
	Visto       int       `gorm:"column:visto;not null;" json:"visto"`
	Reportes_id int64     `gorm:"column:reportes_id;not null;" json:"reporte_id"`
	Usuarios_id int       `gorm:"column:usuarios_id;not null;" json:"usuario_id"`
}

func MensajeCRUD(app *gin.Engine) {
	app.GET("/mensaje/:id", MensajeFetchOne)
	app.GET("/mensaje/", MensajeFetchAll)
	app.POST("/mensaje/", MensajeCreate)
	app.PUT("/mensaje/:id", MensajeUpdate)
	app.DELETE("/mensaje/:id", MensajeRemove)
}

func MensajeFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var mensaje Mensajes
	if err := db.Find(&mensaje, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, mensaje)
	}
}

func MensajeFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var mensajes []Mensajes
	if err := db.Find(&mensajes).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, mensajes)
	}
}

func MensajeCreate(c *gin.Context) {
	var mensaje Mensajes
	e := c.BindJSON(&mensaje)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&mensaje).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, mensaje)
	}
}

func MensajeUpdate(c *gin.Context) {
	var mensaje Mensajes
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&mensaje).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&mensaje)

		db.Save(&mensaje)
		c.JSON(http.StatusOK, mensaje)
	}
}

func MensajeRemove(c *gin.Context) {
	var mensaje Mensajes

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&mensaje).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&mensaje)
		c.JSON(http.StatusOK, mensaje)
	}
}
