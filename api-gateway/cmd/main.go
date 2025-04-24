package main

import (
	"api-gateway/config"
	"api-gateway/grpc"
	"api-gateway/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Env.FrontendURL,
		AllowMethods: "GET, POST, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	shortenerServiceClient := grpc.NewShortenerClient(config.Env)
	redirectionServiceClient := grpc.NewRedirectionClient(config.Env)

	handler := handler.NewHandler(shortenerServiceClient, redirectionServiceClient)

	app.Post("/shortener", handler.CreateShortURL())
	app.Get("/:short_code/origin_url", handler.GetOriginURL())
	app.Get("/:short_code/access_log", handler.GetAccessLog())
	app.Post("/redirect", handler.RedirectHandler())
	app.Get("/urls/:session_id", handler.GetUrls())

	app.Listen(config.Env.Host + ":" + config.Env.Port)
}
