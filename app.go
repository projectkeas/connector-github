package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projectkeas/connector-github/handlers/hmac"
	"github.com/projectkeas/connector-github/handlers/webhooks"
	"github.com/projectkeas/sdks-service/eventPublisher"
	"github.com/projectkeas/sdks-service/server"
)

func main() {
	app := server.New("connector-github")

	app.WithEnvironmentVariableConfiguration("KEAS_")

	app.WithConfigMap("connector-github-cm")
	app.WithRequiredSecret("connector-github-secret")
	app.WithRequiredSecret("ingestion-secret")

	app.ConfigureHandlers(func(f *fiber.App, server *server.Server) {
		f.Route("integrations/github", func(router fiber.Router) {
			router.Post("/webhooks", hmac.NewSha256("X-Hub-Signature-256", server, "github.webhook.token"), webhooks.New(server))
		})
	})

	server := app.Build()

	server.RegisterService(eventPublisher.SERVICE_NAME, eventPublisher.New(server.GetConfiguration()))

	server.Run()
}
