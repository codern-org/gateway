package infrastructure

import (
	"github.com/codern-org/gateway/pkg/route"
	"github.com/gofiber/fiber/v2"
)

func OpenFiberServer() error {
	server := fiber.New()

	route.GatewayRoutes(server)

	return server.Listen(":3000")
}
