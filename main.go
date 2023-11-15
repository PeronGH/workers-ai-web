package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/PeronGH/cf-ai-web-ui/internal/api"
	"github.com/PeronGH/cf-ai-web-ui/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

//go:embed views/*
var views embed.FS

var paths = map[string]string{
	"/":           "Image Generation",
	"/similarity": "Embedding Similarity",
}

func main() {
	err := godotenv.Load()
	if err == nil {
		log.Println("Loaded .env file")
	}

	engine := html.NewFileSystem(http.FS(views), ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/pages/image-generation", fiber.Map{
			"Paths":       paths,
			"CurrentPath": c.Path(),
		}, "views/layouts/main")
	})

	app.Post("/api/run/*", func(c *fiber.Ctx) error {
		model := c.Path()[len("/api/run/"):]
		return api.RunModel(model, c.Body(), c.Response().BodyWriter())
	})

	log.Fatal(app.Listen(":" + utils.GetEnv("PORT", "3000")))
}
