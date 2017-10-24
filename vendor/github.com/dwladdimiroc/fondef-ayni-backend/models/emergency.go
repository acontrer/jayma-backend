package models

import (
	"github.com/gin-gonic/gin"
	//		"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"
	"time"

	"github.com/dwladdimiroc/fondef-ayni-backend/api"
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Emergency struct {
	//	gorm.Model
	Id                  uint             `gorm:"primary_key"`
	CreatedAt           time.Time        `gorm:"column:createAt" json:"createAt"`
	Title               string           `gorm:"column:title;not null;" json:"title"`
	Place_latitude      float64          `gorm:"column:place_latitude;not null;" json:"place_latitude"`
	Place_longitude     float64          `gorm:"column:place_longitude;not null;" json:"place_longitude"`
	Place_radius        float64          `gorm:"column:place_radius;not null;" json:"place_radius"`
	Description         string           `gorm:"column:description;not null;" json:"description"`
	Commune             string           `gorm:"column:commune;not null;" json:"commune"`
	City                string           `gorm:"column:city;not null;" json:"city"`
	Region              string           `gorm:"column:region;not null;" json:"region"`
	Emergency_type_id   int              `gorm:"column:emergency_type_id;not null;" json:"emergency_type_id"`
	Emergency_type      Emergencies_type `gorm:"ForeignKey:Emergency_type_id;AssociationForeignKey:Id"`
	Emergency_status_id int              `gorm:"column:emergency_status_id;not null;" json:"emergency_status_id"`
	Missions            []Mission        `json:"Missions"`
	User_id             int              `gorm:"column:user_id;not null;" json:"user_id"`
	User                User             `gorm:"ForeignKey:User_id;AssociationForeignKey:Id"`
}

func EmergencyCRUD(app *gin.Engine) {
	app.GET("/emergency/:id", EmergencyFetchOne)
	app.GET("/emergency/", EmergencyFetchAll)
	app.POST("/emergency/", EmergencyCreate)
	app.PUT("/emergency/:id", EmergencyUpdate)
	app.DELETE("/emergency/:id", EmergencyRemove)
}

func EmergencyFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var emergency Emergency
	if err := db.Find(&emergency, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Model(&emergency).Related(&emergency.Emergency_type, "Emergency_type_id")
		db.Model(&emergency).Related(&emergency.Missions, "Emergency_id")
		c.JSON(http.StatusOK, emergency)
	}

}

func EmergencyFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var emergencies []Emergency

	if err := db.Find(&emergencies).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		for i := range emergencies {
			db.Model(&emergencies[i]).Related(&emergencies[i].Emergency_type, "Emergency_type_id")
			db.Model(&emergencies[i]).Related(&emergencies[i].Missions, "Emergency_id")
		}
		c.JSON(http.StatusOK, emergencies)
	}

}

func EmergencyCreate(c *gin.Context) {
	var emergency Emergency
	e := c.BindJSON(&emergency)
	utils.Check(e)

	commune, city, region, err := api.EmergencyReverseGeocodingPoint(emergency.Place_latitude, emergency.Place_longitude)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		emergency.Commune = commune
		emergency.City = city
		emergency.Region = region

		db := db.Database()
		defer db.Close()

		if err := db.Create(&emergency).Error; err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusCreated, emergency)
		}
	}
}

func EmergencyUpdate(c *gin.Context) {
	var emergency Emergency
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&emergency).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&emergency)

		db.Save(&emergency)
		c.JSON(200, emergency)
	}
}

func EmergencyRemove(c *gin.Context) {
	var emergency Emergency

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&emergency).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&emergency)
		c.JSON(http.StatusOK, emergency)
	}
}
