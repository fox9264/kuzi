package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL(config map[string]string)(err error){
	DB, err = gorm.Open(config["db_type"], config["db_url"])
	if err != nil {
		return
	}
	DB.LogMode(true)
	return DB.DB().Ping()
}

func Close(){
	DB.Close()
}