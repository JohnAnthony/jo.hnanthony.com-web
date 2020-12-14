//go:generate go-bindata -o static.go static/...

package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"html/template"
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func buildBody() ([]byte, error) {
	bg, err := Asset("static/bg.png")
	if err != nil {
		return []byte{}, err
	}

	templateVars := struct {
		BG string
	}{
		BG: base64.StdEncoding.EncodeToString(bg),
	}

	var ret []byte
	buf := bytes.NewBuffer(ret)
	index, err := Asset("static/index.html.template")
	if err != nil {
		return []byte{}, err
	}
	tmpl, err := template.New("index").Parse(string(index))
	if err != nil {
		return []byte{}, err
	}
	err = tmpl.Execute(buf, templateVars)
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

func buildEtag(body *[]byte) string {
	hasher := sha512.New()
	hasher.Write(*body)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func main() {
	body, err := buildBody()
	if err != nil {
		log.Fatal(err)
	}
	etag := buildEtag(&body)

	attachHeaders := func(c *gin.Context) {
		c.Header("ETag", etag)
		if c.GetHeader("If-None-Match") == etag {
			c.AbortWithStatus(304)
			return
		}
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		attachHeaders(c)
		c.Data(200, "text/html", body)
	})
	r.HEAD("/", func(c *gin.Context) {
		attachHeaders(c)
		c.Data(200, "text/html", []byte{})
	})

	log.Fatal(autotls.Run(r, "jo.hnanthony.com"))
}
