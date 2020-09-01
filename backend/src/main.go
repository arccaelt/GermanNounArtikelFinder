package main

import (
	"dao"
	"routes"
)

func main() {
	dao.InitDAO()
	defer dao.DBORM.Close()

	router := routes.InitRoutes()
	router.Run()
}
