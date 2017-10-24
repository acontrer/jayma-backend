package routes

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

func ProblemCreate(c *gin.Context) {
	if userId, exists := c.Get("userID"); exists {
		var user models.User

		db := db.Database()
		defer db.Close()

		if err := user.UserFetchEmail(userId.(string)); err != nil {
			c.String(http.StatusNotFound, err.Error())
		} else {

			var problem models.Problem
			e := c.BindJSON(&problem)
			utils.Check(e)

			problem.Assertiveness_text = utils.Asertiveness(problem.Description)
			problem.User_id = user.Id

			if err := db.Create(&problem).Error; err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			} else {
				c.JSON(http.StatusCreated, problem)
			}
		}
	}
}
