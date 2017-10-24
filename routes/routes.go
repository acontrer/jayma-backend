package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dwladdimiroc/fondef-jayma-backend/auth"
)

func Setup(app *gin.Engine) {
	authMiddleware := auth.CreateMiddleware()
	app.POST("/login", authMiddleware.LoginHandler)

	authorization := app.Group("/auth")
	authorization.Use(auth.AddPermission(auth.Usuario))
	authorization.Use(authMiddleware.MiddlewareFunc())
	{
		authorization.GET("/permisos", auth.StatusPermission)
		authorization.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
}
