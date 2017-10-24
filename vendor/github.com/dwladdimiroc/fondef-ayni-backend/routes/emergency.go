package routes

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/models"

	"github.com/gin-gonic/gin"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"

	"net/http"
)

func EmergenciasActivas(c *gin.Context) {
	db := db.Database()
	defer db.Close()

	var emergencies []models.Emergency

	if err := db.Where("emergency_status_id = ?", 1).Find(&emergencies).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		for i := range emergencies {
			db.Model(&emergencies[i]).Related(&emergencies[i].Emergency_type, "Emergency_type_id")
			db.Where("mission_status_id = ? OR mission_status_id = ?", 1, 2).Model(&emergencies[i]).Related(&emergencies[i].Missions, "Emergency_id")
		}
		c.JSON(http.StatusOK, emergencies)
	}
}

func EmergenciasArchivadas(c *gin.Context) {
	db := db.Database()
	defer db.Close()

	var emergencies []models.Emergency
	db.Where("emergency_status_id = ?", 2).Find(&emergencies)

	for i := range emergencies {
		db.Model(&emergencies[i]).Related(&emergencies[i].Emergency_type, "Emergency_type_id")
	}

	c.JSON(http.StatusOK, emergencies)
}
