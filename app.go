package main

import (
	"github.com/dwladdimiroc/fondef-jayma-backend/db"
	"github.com/dwladdimiroc/fondef-jayma-backend/models"
	"github.com/dwladdimiroc/fondef-jayma-backend/routes"
	"github.com/dwladdimiroc/fondef-jayma-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig("config/config.yaml")

	db.Setup()

	app := gin.Default()
	app.Use(utils.CorsMiddleware())

	models.Setup(app)
	routes.Setup(app)

	app.Run(":" + utils.Config.Server.Port)
}
