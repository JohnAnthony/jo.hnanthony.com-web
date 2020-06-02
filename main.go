//go:generate go-bindata -o static.go static/...

package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	index, err := Asset("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	bg, err := Asset("static/bg.png")
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.Writer.Write(index)
	})
	r.GET("/bg.png", func(c *gin.Context) {
		c.Writer.Write(bg)
	})
	r.Run()
}
