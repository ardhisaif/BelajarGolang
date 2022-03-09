package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := Cek()
	r.Run(":9000")
}

func Cek() *gin.Engine{
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "cek masuk")
	})
	return r
}