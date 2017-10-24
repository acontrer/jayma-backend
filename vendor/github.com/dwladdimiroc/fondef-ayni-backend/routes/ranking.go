package routes

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"

	"github.com/gin-gonic/gin"

	"net/http"
	"sort"
	"strings"
	//	"fmt"
)

type Abilities struct {
	Abilities []string `json:"abilities" binding:"required"`
}

type UserRanking struct {
	User    models.User `json:"user"`
	Ranking int         `json:"ranking"`
}

type ByRanking []UserRanking

func (ur ByRanking) Len() int           { return len(ur) }
func (ur ByRanking) Swap(i, j int)      { ur[i], ur[j] = ur[j], ur[i] }
func (ur ByRanking) Less(i, j int) bool { return ur[i].Ranking < ur[j].Ranking }

func getAbilities(abilities []models.Ability) []string {
	abilitiesString := make([]string, len(abilities))
	for i := range abilities {
		abilitiesString[i] = abilities[i].Ability
	}

	return abilitiesString
}

func calcularRanking(users []UserRanking, abilitiesMission []string) {
	max := 0

	for i := range users {
		count := 0
		for _, ability := range getAbilities(users[i].User.Volunteer.Abilities) {
			for _, abilityMission := range abilitiesMission {
				if strings.Compare(ability, abilityMission) == 0 {
					count++
				}
			}
		}
		if count > max {
			max = count
		}
		users[i].Ranking = count
	}

	for i := range users {
		if max != 0 {
			valor := float64(users[i].Ranking) / float64(max)
			users[i].Ranking = int(valor * 100)
		} else {
			users[i].Ranking = 0
		}
	}

}

func RankingVoluntarios(c *gin.Context) {
	var abilitiesMission Abilities
	e := c.BindJSON(&abilitiesMission)
	utils.Check(e)

	db := db.Database()
	defer db.Close()

	var volunteers []models.Volunteer
	db.Where("Volunteer_status_id = ?", 1).Find(&volunteers)

	users := make([]UserRanking, len(volunteers))
	for i := range volunteers {
		db.Model(&volunteers[i]).Related(&users[i].User, "User_id")
		users[i].User.Volunteer = volunteers[i]
		db.Model(&users[i].User.Volunteer).Related(&users[i].User.Volunteer.Abilities, "Abilities")
	}

	calcularRanking(users, abilitiesMission.Abilities)

	sort.Sort(ByRanking(users))

	c.JSON(http.StatusOK, users)
}
