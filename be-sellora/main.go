package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi route Gin
	route := gin.Default()

	// Membuat endpoint dengan method GET
	route.Get("/", func(c *gin.Context) {
		
		// Mengirim response dalam format JSON
		c.JSON(200, gin.H {
			"message": "Hello World!",
		})
	})

	// Menjalankan server port 3000
	router.Run(":3000")
}