package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// from https://gist.github.com/schollz/f25d77afc9130b72390748bdbce0d9a3

func main() {
	fmt.Println("Starting server...")
	router := gin.Default()
	router.POST("/many", func(c *gin.Context) { // multiple files
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for _, file := range files {
			log.Println(file.Filename)
			err := c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(c.PostForm("key"))
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.POST("/upload", func(c *gin.Context) { // one file
		// single file
		file, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fileHeader.Filename)

		// err = c.SaveUploadedFile(file, file.Filename)
		f, err := os.Create("goUpload.txt")
		io.Copy(f, file)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileHeader.Filename))
	})
	router.Run(":8080")
}
