package monitor

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	types "github.com/nielchaudhary/compass/internal/monitor/types"
	storage "github.com/nielchaudhary/compass/internal/storage"
	"github.com/nielchaudhary/compass/pkg/constants"
	logger "github.com/nielchaudhary/compass/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func RegisterServiceEndpoints(c *fiber.Ctx) error {
	var (
		req    types.RegisterHealthRequestBody
		zapLog *zap.SugaredLogger
	)

	zapLog = logger.GetLogger("register-service")

	endpointsColl, err := storage.GetCollection("endpoints")
	if err != nil {
		zapLog.Fatal("Error getting endpoints collection, please check", err)

	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body: " + err.Error(),
			"data":    nil,
		})
	}

	if req.Endpoint == "" || req.ServiceName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID, Endpoint & ServiceName are requiredd fields",
			"data":    nil,
		})
	}

	// Validate HTTP method
	method := strings.ToUpper(string(req.Method))
	isValidMethod := slices.Contains(constants.ValidHttpMethods, method)

	if !isValidMethod {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid HTTP method. Must be one of: GET, POST, PUT, PATCH, DELETE",
			"data":    nil,
		})
	}

	req.Method = types.RequestMethod(method)

	_, err = endpointsColl.InsertOne(c.Context(), bson.M{
		"serviceId":   uuid.New(),
		"endpoint":    req.Endpoint,
		"serviceName": req.ServiceName,
		"method":      req.Method,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to insert endpoint in MongoDB due to: " + err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Health endpoint registered successfully",
		"data":    req,
	})
}
