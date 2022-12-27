package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/raexufybyel/janken-go/db"
	"log"
	"net/http"
)

func main() {
	mysql, err := db.NewDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer mysql.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
