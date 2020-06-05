//go:generate go-bindata -o static.go static/...

package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	index, err := Asset("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	bg, err := Asset("static/bg.png")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write(index)
	})
	r.GET("/bg.png", func(c *gin.Context) {
		c.Writer.Write(bg)
	})
	r.Run()
}
