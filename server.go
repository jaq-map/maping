package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("src/views/**")

	router.GET("/", indexHandler)
	router.GET("/index.html", func(c *gin.Context) {
		c.Redirect(302, "/")
	})

	router.GET("/ping", func(c *gin.Context) {
		cmd := exec.Command("bash", "-c", "src/shellscripts/ping.sh")
		out, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.HTML(200, "ping.html", gin.H{"title": "Maping", "message": string(out)})
		}
	})

	router.Run(":80")
}

func indexHandler(c *gin.Context) {
    obj := gin.H{"title": "Maping"}
    c.HTML(200, "index.html", obj)
}