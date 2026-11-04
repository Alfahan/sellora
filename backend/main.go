package main

import "github.com/gin-gonic/gin"

func main() {
	// Inisialisasi router Gin
	router := gin.Default()

	// Membuat endpoint dengan method GET
	router.GET("/", func(c *gin.Context) {

		// Mengirim response dalam format JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Menjalankan server pada port 3000
	router.Run(":3000")
}
