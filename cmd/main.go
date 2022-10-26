package main

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Listen(":3000")
}
