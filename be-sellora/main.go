package main

import (
	"be-sellora/config"
	"be-sellora/database"
	"be-sellora/database/seeders"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load config .env
	config.LoadEnv()

	// inisialisasi database
	database.InitDB()

	// run seeders
	seeders.Seed()

	// Inisialisasi route Gin
	router := gin.Default()

	// Membuat endpoint dengan method GET
	router.GET("/", func(c *gin.Context) {

		// Mengirim response dalam format JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Menjalankan server port 3000
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
