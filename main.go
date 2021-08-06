//go:generate go-bindata -o static.go static/...

package main

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"log"

	fiber "github.com/gofiber/fiber/v2"
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

func main() {
	body, err := buildBody()
	if err != nil {
		log.Fatal(err)
	}

	r := fiber.New()
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Type("html").Send(body)
	})
	log.Fatal(r.Listen(":8080"))
}
