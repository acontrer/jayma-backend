package models

import (
	//		"fmt"

	"time"

	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	//"encoding/json"
	"net/http"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type User struct {
	Id                     int       `gorm:"column:id;primary_key;"`
	First_name             string    `gorm:"column:first_name;not null;" json:"first_name"`
	Last_name              string    `gorm:"column:last_name;not null;" json:"last_name"`
	Birthday               time.Time `gorm:"column:birthday;not null;" json:"birthday"`
	Password               string    `gorm:"column:password;not null;" json:"password"`
	Email                  string    `gorm:"column:email;not null;" json:"email"`
	Contact_phone_number   int       `gorm:"column:contact_phone_number;not null;" json:"contact_phone_number"`
	Emergency_phone_number int       `gorm:"column:emergency_phone_number;not null;" json:"emergency_phone_number"`
	Life_insurance         bool      `gorm:"column:life_insurance;not null;" json:"life_insurance"`
	User_type_id           int       `gorm:"column:user_type_id;not null;" json:"user_type_id"`
	Enabled                bool      `gorm:"column:enabled;not null;" json:"enabled"`
	Volunteer              Volunteer
}

func UserCRUD(app *gin.Engine) {
	app.GET("/user/:id", UserFetchOne)
	app.GET("/userName/:name", UserFetchName)
	app.GET("/user/", UserFetchAll)
	app.POST("/user/", UserCreate)
	app.PUT("/user/:id", UserUpdate)
	app.DELETE("/user/:id", UserRemove)
}

func UserFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var user User

	if err := db.Find(&user, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Model(&user).Related(&user.Volunteer, "Users_id")

		db.Model(&user.Volunteer).Related(&user.Volunteer.Abilities, "Abilities")

		c.JSON(http.StatusOK, user)
	}
}

func UserFetchName(c *gin.Context) {
	name := c.Param("name")

	db := db.Database()
	defer db.Close()

	var user User
	db.Where("first_name = ?", name).First(&user)
	db.Model(&user).Related(&user.Volunteer, "Users_id")
	db.Model(&user.Volunteer).Related(&user.Volunteer.Abilities, "Abilities")

	c.JSON(http.StatusOK, user)
}

func UserFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		for i := range users {
			db.Model(&users[i]).Related(&users[i].Volunteer, "Users_id")
			db.Model(&users[i].Volunteer).Related(&users[i].Volunteer.Abilities, "Abilities")
		}
		c.JSON(http.StatusOK, users)
	}
}

func UserCreate(c *gin.Context) {
	var user User
	e := c.BindJSON(&user)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

func UserUpdate(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&user)
		db.Save(&user)
		c.JSON(200, user)
	}
}

func UserRemove(c *gin.Context) {
	var user User

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {

	} else {
		db.Delete(&user)
		c.JSON(http.StatusOK, user)
	}
}

func UserId(email string) int {
	db := db.Database()
	defer db.Close()

	var user User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return 0
	} else {
		return user.Id
	}
}

func (user *User) UserFetchEmail(email string) error {
	db := db.Database()
	defer db.Close()

	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	} else {
		db.Model(&user).Related(&user.Volunteer, "Users_id")
		db.Model(&user.Volunteer).Related(&user.Volunteer.Abilities, "Abilities")
		return nil
	}
}
