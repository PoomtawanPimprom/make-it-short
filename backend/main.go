package main

import (
	"github/PoomtawanPimprom/make-it-short/database"
	"github/PoomtawanPimprom/make-it-short/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	routes.UrlRoutes(r)
	r.Run(":8081")
}
