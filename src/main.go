package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")

	db, err := sql.Open("mysql", "devuser:devpass@tcp(db:3306)/development")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var dbVersion string
	db.QueryRow("SELECT VERSION()").Scan(&dbVersion)
	fmt.Println("Connected to: ", dbVersion)

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":     "Go for it!!",
			"dbVersion": dbVersion,
		})
	})
	engine.Run(":3000")
}
