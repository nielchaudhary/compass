package monitor

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	types "github.com/nielchaudhary/compass/internal/monitor/types"
)

func RegisterServiceEndpoints(c *fiber.Ctx) error {
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
	method := strings.ToUpper(string(req.Method))
	isValidMethod := slices.Contains(validMethods, method)

	if !isValidMethod {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid HTTP method. Must be one of: GET, POST, PUT, PATCH, DELETE",
			"data":    nil,
		})
	}

	req.Method = types.RequestMethod(method)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Health endpoint registered successfully",
		"data":    req,
	})
}
