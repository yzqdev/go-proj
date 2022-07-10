package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gookit/color"
	"gofiber/controller"
)

func createRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about")
	})
	app.Get("/users", controller.GetUsers)
	color.Redln("routes")
	app.Get("/view", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	app.Static("/public", "./public")
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	createRouteMap(app)

}
func createRouteMap(engine *fiber.App) {
	routes := engine.Stack()
	for _, route := range routes {
		for _, r := range route {
			color.Redln("[debug]", r.Method, r.Path, r.Params)
		}
	}

}
