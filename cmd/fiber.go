package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	engine := html.New("./public",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(recover.New())

	app.Static("/","./public")

	app.Get("/crash_all", func(c *fiber.Ctx) error {
		panic("This panic is caught by fiber")
	})

	api := app.Group("/api")

	api.Get("/*", func (c *fiber.Ctx) error {
		value := c.Params("*")

		switch value {
			case "auth":
			return c.SendString("PATH: " + value)
			case "logout":
			return c.SendString("PATH: " + value)
			case "send_mail":
			return c.SendString("PATH: " + value)
			case "apchihbva":
			break
			case "show":
			return c.Render("main", fiber.Map{
				"Title": "Something",
			})
			default:
			// return c.SendString("Non Valid Path")
			return c.Next()
		}

		return c.SendString("Vse poshlo po pizde")
	})

	api.Use([]string{"abc_","abc2_"},func (c *fiber.Ctx) error {
		return c.SendString("abc_code")
	})

	api.Get("/product/color::color/size::size/:select<bool>", func(c *fiber.Ctx) error {
		if (c.Params("select") == "true") {
			return c.SendString("Quality - color " + c.Params("color") + " - size " + c.Params("size"))
		}
		return c.SendString("Not selected product")
	})

	api.Get("/:random.txt", func (c *fiber.Ctx) error {
		return c.SendString("no text for now like a " + c.Params("random"))
	})

	app.Post("/", func (c *fiber.Ctx) error {
		return c.SendString("POST request")
	})

	log.Fatal(app.Listen(":3000"))
}
