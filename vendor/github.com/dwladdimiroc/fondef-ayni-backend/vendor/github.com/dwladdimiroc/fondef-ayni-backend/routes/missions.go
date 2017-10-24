package routes

import (
	"github.com/dwladdimiroc/fondef-ayni-backend/db"
	"github.com/dwladdimiroc/fondef-ayni-backend/models"
	"github.com/dwladdimiroc/fondef-ayni-backend/utils"

	"github.com/gin-gonic/gin"

	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func MisionesSubirArchivo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("idMission"))
	utils.Check(err)

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	src, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer src.Close()

	err = utils.CreateFolder(utils.Config.Files.Uri + c.Param("idMission"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dir := filepath.Join(utils.Config.Files.Uri, c.Param("idMission"))
	name := utils.CreateFilename(dir, file.Filename)
	filename := filepath.Join(dir, name)

	dst, err := os.Create(filename)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer dst.Close()

	io.Copy(dst, src)

	var fileDb models.File
	fileDb.File = name
	fileDb.Mission_id = id

	db := db.Database()
	defer db.Close()

	if err := db.Create(&fileDb).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusCreated, fileDb)
	}
}

func MisionesActivas(c *gin.Context) {
	db := db.Database()
	defer db.Close()

	var missions []models.Mission
	if err := db.Where("mission_status_id = ? OR mission_status_id = ?", 1, 2).Find(&missions).Error; err != nil {
		c.String(http.StatusNotFound, err.Error())
	} else {
		var missionsActive []models.Mission
		for i := range missions {
			db.Model(&missions[i]).Related(&missions[i].User, "User_id")
			db.Model(&missions[i]).Related(&missions[i].Emergency, "Emergency_id")
			if missions[i].Emergency.Emergency_status_id == 1 {
				missionsActive = append(missionsActive, missions[i])
			}
		}

		c.JSON(http.StatusOK, missionsActive)
	}
}

func MisionesFinalizadas(c *gin.Context) {
	db := db.Database()
	defer db.Close()

	var missions []models.Mission
	db.Where("mission_status_id = ?", 2).Find(&missions)

	for i := range missions {
		db.Model(&missions[i]).Related(&missions[i].User, "User_id")
		db.Model(&missions[i]).Related(&missions[i].Emergency, "Emergency_id")
	}

	c.JSON(http.StatusOK, missions)
}

func MisionesArchivadas(c *gin.Context) {
	db := db.Database()
	defer db.Close()

	var missions []models.Mission
	db.Where("mission_status_id = ?", 3).Find(&missions)

	for i := range missions {
		db.Model(&missions[i]).Related(&missions[i].User, "User_id")
		db.Model(&missions[i]).Related(&missions[i].Emergency, "Emergency_id")
	}

	c.JSON(http.StatusOK, missions)
}

func MisionesPorcentajeAvance(c *gin.Context) {
	id := c.Param("idMission")

	db := db.Database()
	defer db.Close()

	var total int
	db.Model(&models.History_missions{}).Where("mission_id = ?", id).Count(&total)

	var finish int
	db.Model(&models.History_missions{}).Where("mission_id = ? AND history_mission_state_id = ?", id, 4).Count(&finish)

	type Response struct {
		Porcentaje int `json:"porcentaje"`
	}
	var resp Response
	var valor float64
	if total != 0 {
		valor = float64(finish) / float64(total)
	} else {
		valor = 0
	}

	resp.Porcentaje = int(valor * 100)

	c.JSON(http.StatusOK, resp)
}
