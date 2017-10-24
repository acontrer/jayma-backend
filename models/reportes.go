package models

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"time"

	"github.com/dwladdimiroc/fondef-jayma-backend/db"
	"github.com/dwladdimiroc/fondef-jayma-backend/utils"
)

type Reportes struct {
	Id            int64     `gorm:"column:id;not null;" json:"id"`
	CreatedAt      time.Time `gorm:"column:fecha_creacion;not null;" json:"fecha_creacion"`
	Estado        bool      `gorm:"column:estado;not null;" json:"estado"`
	Posicion_lat  float64   `gorm:"column:posicion_lat;not null;" json:"posicion_lat"`
	Posicion_long float64   `gorm:"column:posicion_long;not null;" json:"posicion_long"`
	Usuarios_id   int       `gorm:"column:usuarios_id;not null;" json:"usuario_id"`
	//Estados       []Estados `gorm:"many2many"`
}

func ReporteCRUD(app *gin.Engine) {
	app.GET("/reporte/:id", ReporteFetchOne)
	app.GET("/reporte/", ReporteFetchAll)
	app.POST("/reporte/", ReporteCreate)
	app.PUT("/reporte/:id", ReporteUpdate)
	app.DELETE("/reporte/:id", ReporteRemove)
}

func ReporteFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var reporte Reportes
	if err := db.Find(&reporte, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, reporte)
	}
}

func ReporteFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var reportes []Reportes
	if err := db.Find(&reportes).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, reportes)
	}
}

func ReporteCreate(c *gin.Context) {
	var reporte Reportes
	e := c.BindJSON(&reporte)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&reporte).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, reporte)
	}
}

func ReporteUpdate(c *gin.Context) {
	var reporte Reportes
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&reporte).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&reporte)

		db.Save(&reporte)
		c.JSON(http.StatusOK, reporte)
	}
}

func ReporteRemove(c *gin.Context) {
	var reporte Reportes

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&reporte).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&reporte)
		c.JSON(http.StatusOK, reporte)
	}
}
