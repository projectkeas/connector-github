package webhooks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projectkeas/connector-github/services/eventBuilder"
	"github.com/projectkeas/sdks-service/eventPublisher"
	"github.com/projectkeas/sdks-service/server"
)

func New(server *server.Server) func(c *fiber.Ctx) error {
	// publisher setup
	nc, err := server.GetService(eventPublisher.SERVICE_NAME)
	if err != nil {
		panic(err)
	}
	client := (*nc).(eventPublisher.EventPublisherService)

	return func(context *fiber.Ctx) error {
		context.Accepts("application/json")

		errorResult := map[string]interface{}{
			"message": "An error occurred whilst processing your request",
		}

		eventName := context.Get("X-GitHub-Event")
		uuid := context.Get("X-GitHub-Delivery")
		if uuid == "" || eventName == "" {
			context.SendStatus(fiber.StatusBadRequest)
			return nil
		}

		if eventName == "ping" {
			context.SendStatus(fiber.StatusNoContent)
			return nil
		}
		payload := map[string]interface{}{}
		err := context.BodyParser(&payload)
		if err != nil {
			errorResult["reason"] = "Unable to parse request body"
			return context.Status(fiber.StatusBadRequest).JSON(errorResult)
		}

		cloudEvent, validationErr := eventBuilder.NewCloudEventFromWebhook(payload, eventName, uuid)
		if validationErr != nil {
			errorResult["reason"] = "Unable to validate as a cloud event"
			errors := []map[string]string{}
			for key, value := range validationErr {
				errors = append(errors, map[string]string{
					"attribute": key,
					"error":     value.Error(),
				})
			}
			errorResult["errors"] = errors
			return context.Status(fiber.StatusBadRequest).JSON(errorResult)
		}

		if !client.Publish(cloudEvent) {
			errorResult["reason"] = "publish"
			return context.Status(fiber.StatusInternalServerError).JSON(errorResult)
		}

		context.Status(fiber.StatusAccepted).JSON(cloudEvent)
		return nil
	}
}
