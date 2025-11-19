package monitor

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	types "github.com/nielchaudhary/compass/internal/monitor/types"
)

func RegisterHealthEndpoint(c *fiber.Ctx) error {
	var req types.RegisterHealthRequestBody

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
	}

	if req.ID == "" || req.Endpoint == "" || req.ServiceName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID, Endpoint & ServiceName are required fields",
			"data":    nil,
		})
	}

	// Validate HTTP method
	validMethods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	method := strings.ToUpper(req.Method)
	isValidMethod := false
	for _, validMethod := range validMethods {
		if method == validMethod {
			isValidMethod = true
			break
		}
	}

	if !isValidMethod {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid HTTP method. Must be one of: GET, POST, PUT, PATCH, DELETE",
			"data":    nil,
		})
	}

	req.Method = method

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Health endpoint registered successfully",
		"data":    req,
	})
}
