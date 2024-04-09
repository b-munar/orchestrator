package main

import (
	"fiber-orchestrator/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowMethods: "GET,POST",
	}))

	app.Use(logger.New())

	app.Get("/orchestrator/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("ping")
	})

	app.Post("/orchestrator/register/sportman", handlers.RegisterSportman)
	app.Post("/orchestrator/register/partner", handlers.RegisterPartner)
	app.Post("/orchestrator/login", handlers.Login)

	app.Listen(":80")
}
