package main

import (
	"sellora/backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config .env
	config.LoadEnv()

	// Inisialisasi Gin
	router := gin.Default()

	// Membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {
		// return response JSON
		c.JSON(200, gin.H{"message": "Hello World!!"})
	})

	// Menjalankan server berdasarkan config
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))

}
