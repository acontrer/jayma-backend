package routes

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
)

func CountAnswer(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	type Count struct {
		Count int `json:"count"`
	}

	var count Count
	if err := db.Where("problem_id = ?", id).Model(&models.Answer{}).Count(&count.Count).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, count)
	}
}
