package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/dwladdimiroc/fondef-ayni-backend/utils"
)

var (
	rdbms, user, pass, ip, port, database, sslmode string
)

func Setup() {
	rdbms = utils.Config.Database.Rdbms
	user = utils.Config.Database.User
	pass = utils.Config.Database.Pass
	ip = utils.Config.Database.Ip
	port = utils.Config.Database.Port
	database = utils.Config.Database.Name
	sslmode = utils.Config.Database.Sslmode
}

func Database() *gorm.DB {
	db, e := gorm.Open(rdbms, rdbms+"://"+user+":"+pass+"@"+ip+":"+port+"/"+database+"?sslmode="+sslmode)
	utils.Check(e)
	return db
}
