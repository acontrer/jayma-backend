package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	"net/http"
	"time"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Answer struct {
	//gorm.Model
	Id         uint      `gorm:"primary_key"`
	CreatedAt  time.Time `gorm:"column:createAt" json:"createAt"`
	Answer     string    `gorm:"column:answer;not null;" json:"answer"`
	Problem_id int       `gorm:"column:problem_id;not null;" json:"problem_id"`
	User_id    int       `gorm:"column:user_id;not null;" json:"user_id"`
	User       User      `gorm:"ForeignKey:User_id;AssociationForeignKey:Id"`
}

func AnswerCRUD(app *gin.Engine) {
	app.GET("/answer/:id", AnswerFetchOne)
	app.GET("/answer/", AnswerFetchAll)
	app.POST("/answer/", AnswerCreate)
	app.PUT("/answer/:id", AnswerUpdate)
	app.DELETE("/answer/:id", AnswerRemove)
}

func AnswerFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var answer Answer
	if err := db.Find(&answer, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Model(&answer).Related(&answer.User, "User_id")
		c.JSON(http.StatusOK, answer)
	}
}

func AnswerFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var answers []Answer
	db.Find(&answers)

	c.JSON(http.StatusOK, answers)
}

func AnswerCreate(c *gin.Context) {
	var answer Answer
	e := c.BindJSON(&answer)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&answer).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, answer)
	}
}

func AnswerUpdate(c *gin.Context) {
	var answer Answer
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&answer).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&answer)

		db.Save(&answer)
		c.JSON(200, answer)
	}
}

func AnswerRemove(c *gin.Context) {
	var answer Answer

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&answer).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&answer)
		c.JSON(http.StatusOK, answer)
	}
}
