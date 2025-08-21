package controllers

import (
	"github/PoomtawanPimprom/make-it-short/database"
	"github/PoomtawanPimprom/make-it-short/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateShortURL(c *gin.Context) {
	var body struct {
		URL string `json:"url" binding:"required,url"`
	}

	if c.Bind(&body) != nil || body.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	if database.DB.Where("original_url = ?", body.URL).First(&models.Url{}).RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "URL already shortened"})
		return
	}

	shortCode := generateCode(6)

	url := models.Url{
		OriginalURL: body.URL,
		ShortCode:   shortCode,
	}

	database.DB.Create(&url)

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8081/" + shortCode,
	})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	var url models.Url
	if err := database.DB.First(&url, "short_code = ?", code).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// update clicks
	url.Clicks++
	database.DB.Save(&url)

	c.Redirect(http.StatusFound, url.OriginalURL)
}

func generateCode(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
