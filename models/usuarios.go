package models

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"time"

	"github.com/dwladdimiroc/fondef-jayma-backend/db"
	"github.com/dwladdimiroc/fondef-jayma-backend/utils"
)

type Usuarios struct {
	Id               int       `gorm:"column:id;not null;"`
	Mail             string    `gorm:"column:mail;not null;"`
	Pass             string    `gorm:"column:pass;not null;"`
	Nombre_primero   string    `gorm:"column:nombre_primero;not null;"`
	Nombre_segundo   string    `gorm:"column:nombre_segundo;not null;"`
	Apellido_paterno string    `gorm:"column:apellido_paterno;not null;"`
	Apellido_materno string    `gorm:"column:apellido_materno;"`
	Fecha_nacimiento time.Time `gorm:"column:fecha_nacimiento;not null;"`
	Telefono         string    `gorm:"column:telefono;not null;"`
	Fb_Usuario       string    `gorm:"column:fb_usuario;not null;"`
}

func UsuarioCRUD(app *gin.Engine) {
	app.GET("/usuario/:id", UsuarioFetchOne)
	app.GET("/usuario/", UsuarioFetchAll)
	app.POST("/usuario/", UsuarioCreate)
	app.PUT("/usuario/:id", UsuarioUpdate)
	app.DELETE("/usuario/:id", UsuarioRemove)
}

func UsuarioFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var usuario Usuarios
	if err := db.Find(&usuario, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, usuario)
	}
}

func UsuarioFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var usuarios []Usuarios
	if err := db.Find(&usuarios).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, usuarios)
	}
}

func UsuarioCreate(c *gin.Context) {
	var usuario Usuarios
	e := c.BindJSON(&usuario)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&usuario).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, usuario)
	}
}

func UsuarioUpdate(c *gin.Context) {
	var usuario Usuarios
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&usuario).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&usuario)

		db.Save(&usuario)
		c.JSON(http.StatusOK, usuario)
	}
}

func UsuarioRemove(c *gin.Context) {
	var usuario Usuarios

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&usuario).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&usuario)
		c.JSON(http.StatusOK, usuario)
	}
}
