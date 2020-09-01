package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBORM *gorm.DB

func InitDAO() {
	db, err := gorm.Open("postgres", "no credentials will be leaked :)")

	if err != nil {
		println(err.Error())
		panic("Can't connect to the database")
	}

	DBORM = db
}
