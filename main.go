package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/PeronGH/workers-ai-web/internal/api"
	"github.com/PeronGH/workers-ai-web/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

//go:embed views/*
var views embed.FS

var paths = fiber.Map{
	"/":              "Image Generation",
	"/similarity":    "Embedding Similarity",
	"/transcription": "Transcription",
}

func main() {
	err := godotenv.Load()
	if err == nil {
		log.Println("Loaded .env file")
	}

	if !utils.HasEnv("CF_API_TOKEN") || !utils.HasEnv("CF_ACCOUNT_ID") {
		log.Fatal("Missing mandatory environment variables")
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

	app.Get("/similarity", func(c *fiber.Ctx) error {
		return c.Render("views/pages/embedding-similarity", fiber.Map{
			"Paths":       paths,
			"CurrentPath": c.Path(),
		}, "views/layouts/main")
	})

	app.Get("/transcription", func(c *fiber.Ctx) error {
		return c.Render("views/pages/transcription", fiber.Map{
			"Paths":       paths,
			"CurrentPath": c.Path(),
		}, "views/layouts/main")
	})

	app.Post("/api/run/*", func(c *fiber.Ctx) error {
		model := c.Path()[len("/api/run/"):]
		return api.RunModel(model, c.Body(), c.Response().BodyWriter())
	})

	app.Get("/api/run/*", func(c *fiber.Ctx) error {
		model := c.Path()[len("/api/run/"):]
		input := c.Queries()

		return api.RunModel(model, input, c.Response().BodyWriter())
	})

	log.Fatal(app.Listen(":" + utils.GetEnv("PORT", "3000")))
}
