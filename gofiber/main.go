package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/gookit/color"
	"log"
	"os"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/dont_compress"
		},
		Level: compress.LevelBestSpeed, // 1
	}))

	createRoute(app)
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")

	err := os.WriteFile("./routes.txt", data, 0777)
	if err != nil {
		color.Redln(err.Error())
		panic(err.Error())

	}
	color.Cyanln("created routes.txt")
	log.Fatal(app.Listen(":3000"))
}
