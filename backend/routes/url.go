package routes

import (
	"github/PoomtawanPimprom/make-it-short/controllers"
	"github.com/gin-gonic/gin"
)

func UrlRoutes(r *gin.Engine) {
	r.POST("/shorten", controllers.CreateShortURL)
	r.GET("/:code", controllers.RedirectURL)
}