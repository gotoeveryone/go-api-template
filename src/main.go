package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	var err error
	time.Local, err = time.LoadLocation(getEnv("TZ", "Asia/Tokyo"))
	if err != nil {
		logrus.Error(fmt.Sprintf("Get location error: %s", err))
		// continue with default timezone.
	}

	r := gin.Default()
	r.HandleMethodNotAllowed = true

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Hello, World!",
		})
	})

	host := getEnv("APP_HOST", "0.0.0.0")
	port := getEnv("APP_PORT", "8080")
	if err := r.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
