package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"testerbackend/models"
	"testerbackend/services"
)

func RunLoadTest(c *fiber.Ctx) error {
	var testReq models.LoadTestRequest
	if err := c.BodyParser(&testReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	fmt.Println(testReq.Duration);
	fmt.Println(testReq.Host);
	fmt.Println(testReq.Requests[0]);

	result, err := services.ExecuteLoadTest(testReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process load test"})
	}

	return c.JSON(result)
}
