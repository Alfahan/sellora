package main

import (
	"sellora/backend/config"
	"sellora/backend/database"
	"sellora/backend/database/seeders"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config .env
	config.LoadEnv()

	// inisialisasi database
	database.InitDB()

	//run seeders
	seeders.Seed()

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
