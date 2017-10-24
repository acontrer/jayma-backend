package models

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	ContactoCRUD(app)
	EstadoCRUD(app)
	MensajeCRUD(app)
	ReporteCRUD(app)
	UsuarioCRUD(app)
}
