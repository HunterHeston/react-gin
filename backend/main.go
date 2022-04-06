package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hunterheston/experiments/react-gin/backend/handlers"
)

func main() {
	r := gin.Default()

	// allow the frontend to make requests
	r.Use(cors.New(cors.Config{
		// address of the frontend when running `npm run dev`
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.POST("/file-upload", handlers.FileUpload)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
