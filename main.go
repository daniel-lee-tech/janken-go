package main

import (
	"github.com/gin-gonic/gin"
	"log"
)
import "net/http"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "dev_user:dev_password@tcp(127.0.0.1:3306)/janken_dev")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
