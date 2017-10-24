package models

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"

	"net/http"
	"time"

	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

type Problem struct {
	//	gorm.Model
	Id                 uint      `gorm:"primary_key"`
	CreatedAt          time.Time `gorm:"column:createAt" json:"createAt"`
	Title              string    `gorm:"column:title;not null;" json:"title"`
	Description        string    `gorm:"column:description;not null;" json:"description"`
	Status             int       `gorm:"column:status;not null;" json:"status"`
	Assertiveness_text float64   `gorm:"column:assertiveness_text;not null;" json:"assertiveness_text"`
	Mission_id         int       `gorm:"column:mission_id;not null;" json:"mission_id"`
	Mission            Mission   `gorm:"ForeignKey:Mission_id;AssociationForeignKey:Id"`
	User_id            int       `gorm:"column:user_id;not null;" json:"user_id"`
	User               User      `gorm:"ForeignKey:User_id;AssociationForeignKey:Id"`
	Answer             []Answer
}

func ProblemCRUD(app *gin.Engine) {
	app.GET("/problem/:id", ProblemFetchOne)
	app.GET("/problem/", ProblemFetchAll)
	app.POST("/problem/", ProblemCreate)
	app.PUT("/problem/:id", ProblemUpdate)
	app.DELETE("/problem/:id", ProblemRemove)
}

func ProblemFetchOne(c *gin.Context) {
	id := c.Param("id")

	db := db.Database()
	defer db.Close()

	var problem Problem
	if err := db.Find(&problem, id).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Model(&problem).Related(&problem.Answer, "Problem_id")
		db.Model(&problem).Related(&problem.User, "User_id")
		db.Model(&problem).Related(&problem.Mission, "Mission_id")

		for i := range problem.Answer {
			db.Model(&problem).Related(&problem.Answer[i], "Problem_id")
			db.Model(&problem.Answer[i]).Related(&problem.Answer[i].User, "User_id")
		}

		c.JSON(http.StatusOK, problem)
	}
}

func ProblemFetchAll(c *gin.Context) {

	db := db.Database()
	defer db.Close()

	var problems []Problem
	db.Find(&problems)

	for i := range problems {
		db.Model(&problems[i]).Related(&problems[i].Answer, "Problem_id")
		db.Model(&problems[i]).Related(&problems[i].User, "User_id")
		db.Model(&problems[i]).Related(&problems[i].Mission, "Mission_id")
	}

	c.JSON(http.StatusOK, problems)
}

func ProblemCreate(c *gin.Context) {
	var problem Problem
	e := c.BindJSON(&problem)
	utils.Check(e)

	problem.Assertiveness_text = utils.Asertiveness(problem.Description)

	db := db.Database()
	defer db.Close()

	if err := db.Create(&problem).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, problem)
	}
}

func ProblemUpdate(c *gin.Context) {
	var problem Problem
	id := c.Params.ByName("id")

	db := db.Database()
	defer db.Close()

	if err := db.Where("id = ?", id).First(&problem).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		c.BindJSON(&problem)

		db.Save(&problem)
		c.JSON(200, problem)
	}
}

func ProblemRemove(c *gin.Context) {
	var problem Problem

	db := db.Database()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&problem).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		db.Delete(&problem)
		c.JSON(http.StatusOK, problem)
	}
}
