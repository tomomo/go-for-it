package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// engine := gin.Default()
	// engine.LoadHTMLGlob("templates/*.html")

	// engine.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"title": "Go for it!",
	// 	})
	// })
	// engine.Run(":3000")

	router := gin.Default()
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "hello",
				})
			})
		}
	}
	router.Run(":3000")

}
