package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Name string `json:"nAmE"`
			Id   int    `json:"id"`
		}{
			Name: "Tricks",
			Id: func() int {
				value, _ := strconv.Atoi(c.Param("id"))
				return value
			}(),
		})
	})

	// Listen and serve on 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
