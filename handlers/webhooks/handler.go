package webhooks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projectkeas/connector-github/services/dataMapper"
	"github.com/projectkeas/sdks-service/server"
)

func New(server *server.Server) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		event := c.Get("X-GitHub-Event")
		uuid := c.Get("X-GitHub-Delivery")
		if uuid == "" || event == "" {
			c.SendStatus(fiber.StatusBadRequest)
			return nil
		}

		if event == "ping" || event == "projects_v2_item" {
			// We explicitly don't support the projects v2 item whilst its in beta (the hook is :puke:)
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}

		src := map[string]interface{}{}
		c.BodyParser(&src)
		dest := dataMapper.MapWebhook(src, event)

		c.SendStatus(fiber.StatusAccepted)
		c.JSON(dest) // returning the mapped properties can be useful for debugging purposes
		return nil
	}
}
