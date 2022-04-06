package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("error parsing form: ", err)
	}
	t := form.Value["test"]
	fmt.Println("test: ", t)

	file, ok := form.File["file"]
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println("Failed to parse file: ", err)
		return
	}
	fmt.Println("File: ", file)
}
