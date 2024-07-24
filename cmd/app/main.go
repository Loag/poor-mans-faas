package main

import (
	event "faas/internal/events"
	"faas/internal/runner"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

const base_path = "/opt/"

func main() {
	router := gin.Default()

	run := runner.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		filePath := base_path + "to_run"
		outFile, err := os.Create(filePath)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		err = os.Chmod(filePath, 755)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "File uploaded successfully"})
	})

	// this is the main handler for all of our functions
	router.NoRoute(func(c *gin.Context) {
		defer c.Request.Body.Close()

		req, err := event.FromRequest(c.Request)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		res := run.Run(req)

		c.JSON(200, gin.H{"message": res.Message, "body": res.Body})
	})

	router.Run(":3000")
}
