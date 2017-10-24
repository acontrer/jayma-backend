package routes

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/auth"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	authMiddleware := auth.CreateMiddleware()

	app.POST("/login", authMiddleware.LoginHandler)

	authorization := app.Group("/auth")
	authorization.Use(auth.AddPermission(auth.Admin, auth.Coordinator, auth.Volunteer))
	authorization.Use(authMiddleware.MiddlewareFunc())
	{
		authorization.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	mission := app.Group("/mision")
	{
		mission.GET("/activas/", MisionesActivas)
		mission.GET("/archivadas/", MisionesArchivadas)
		mission.POST("/archivo/:idMission", MisionesSubirArchivo)
		mission.GET("/avance/:idMission", MisionesPorcentajeAvance)
	}

	emergency := app.Group("/emergencia")
	{
		emergency.GET("/activas/", EmergenciasActivas)
		emergency.GET("/archivadas/", EmergenciasArchivadas)
	}

	voluntarios := app.Group("/voluntarios")
	{
		voluntarios.GET("/disponibles/", VoluntariosDisponibles)
		voluntarios.POST("/ranking/", RankingVoluntarios)
		voluntarios.POST("/interes/", VoluntarioInteres)
	}

	mobile := app.Group("/")
	mobile.Use(auth.AddPermission(auth.Volunteer))
	mobile.Use(authMiddleware.MiddlewareFunc())
	{
		mission := mobile.Group("/mision")
		{
			mission.GET("/invitaciones/", InvitacionesPorUsuario)
			mission.GET("/activa/", MisionActivaUsuario)
		}

		user := mobile.Group("/usuario")
		{
			user.GET("/informacion/", FetchUserInformacion)
			user.PUT("/informacion/", EditUserInformacion)
			user.PUT("/token/", EditTokenUser)
		}

		historyMission := mobile.Group("/historia")
		{
			historyMission.GET("/:mission", ObtenerEstadoUsuario)
			historyMission.PUT("/:mission", CambiarEstadoUsuario)
		}

		problem := mobile.Group("/problema")
		{
			problem.POST("/", ProblemCreate)
		}

		answer := mobile.Group("/respuesta")
		{
			answer.GET("/cant/:id", CountAnswer)
		}
	}
}
