package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/article", handleWordArticle)

	return router
}
