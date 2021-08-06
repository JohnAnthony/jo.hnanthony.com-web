package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"html/template"
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed static/bg.png
var bg []byte

//go:embed static/index.html.template
var indexTemplate string

func buildBody() ([]byte, error) {
	templateVars := struct {
		BG string
	}{
		BG: base64.StdEncoding.EncodeToString(bg),
	}

	var ret []byte
	buf := bytes.NewBuffer(ret)
	tmpl, err := template.New("index").Parse(indexTemplate)
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

	r := fiber.New(fiber.Config{
		GETOnly:          true,
		DisableKeepalive: true,
	})
	r.Use(logger.New())
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Type("html").Send(body)
	})
	log.Fatal(r.Listen(":8080"))
}
