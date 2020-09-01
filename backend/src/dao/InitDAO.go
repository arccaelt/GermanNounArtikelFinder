package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBORM *gorm.DB

func InitDAO() {
	db, err := gorm.Open("postgres", "sslmode=disable host=localhost port=5432 user=postgres dbname=words password=...parola...")

	if err != nil {
		println(err.Error())
		panic("Can't connect to the database")
	}

	DBORM = db
}
