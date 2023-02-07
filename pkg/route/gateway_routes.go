package route

import "github.com/gofiber/fiber/v2"

func GatewayRoutes(server *fiber.App) {
	routes := server.Group("/")

	routes.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello from Codern API",
		})
	})
}
