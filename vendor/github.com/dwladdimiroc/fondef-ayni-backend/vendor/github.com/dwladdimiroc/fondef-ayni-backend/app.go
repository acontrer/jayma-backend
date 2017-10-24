package main

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
	"github.com/dwladdimiroc/fondef-ayni-backend/routes"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig("config/config.yaml")

	db.Setup()

	app := gin.Default()
	app.Use(cors.Default())

	routes.Setup(app)
	models.Setup(app)

	app.Static("/files", "./files")
	app.Run(":" + utils.Config.Server.Port)
}
