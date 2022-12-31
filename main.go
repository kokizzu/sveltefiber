package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"

	"sveltefiber/business"
	"sveltefiber/model"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	const indexFile = "svelte/index.html"
	template, err := os.ReadFile(indexFile)
	L.PanicIf(err, `os.ReadFile: `+indexFile)

	app.Get("/", func(c *fiber.Ctx) error {
		tpl := strings.ReplaceAll(string(template), `/*! title */`, `Hewlloo World`)
		tpl = strings.ReplaceAll(tpl, `[/* raw_data */]`, X.ToJson(model.ListNames()))
		c.Set("Content-type", "text/html")
		return c.SendString(tpl)
	})
	app.Post("/api/createName", func(c *fiber.Ctx) error {
		in := business.CreateNameIn{}
		if err := c.BodyParser(&in); err != nil {
			return err
		}
		out := business.CreateName(&in)
		L.Describe(in, out)
		return c.JSON(out)
	})
	app.Post("/api/deleteName", func(c *fiber.Ctx) error {
		in := business.DeleteNameIn{}
		if err := c.BodyParser(&in); err != nil {
			return err
		}
		out := business.DeleteName(&in)
		L.Describe(in, out)
		return c.JSON(out)
	})

	log.Fatal(app.Listen(":3001"))
}
