//go:generate go-bindata -o static.go static/...

package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
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

	store := persistence.NewInMemoryStore(999999 * time.Hour)

	r := gin.Default()
	r.GET("/", cache.CachePage(store, 999999*time.Hour, func(c *gin.Context) {
		c.Writer.Write(index)
	}))
	r.GET("/bg.png", cache.CachePage(store, 999999*time.Hour, func(c *gin.Context) {
		c.Writer.Write(bg)
	}))
	r.Run()
}
