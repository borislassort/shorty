package main

// Need CGO_ENABLED=1

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var urlMap = make(map[string]string)
var rng *rand.Rand

func main() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	r := gin.Default()
	r.POST("/shorten", shortURL)
	r.GET("/:shortURL", redirectURL)
	r.Run(":8080")
}

func shortURL(c *gin.Context) {
	longURL := c.Query("url")
	shortURL := genShortURL()
	urlMap[shortURL] = longURL
	c.JSON(http.StatusOK, gin.H{"longURL": longURL, "shortURL": shortURL})
}

func redirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	longURL, ok := urlMap[shortURL]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Short URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, longURL)
}

func genShortURL() string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = charset[rng.Intn(len(charset))]
	}
	return string(shortURL)
}
