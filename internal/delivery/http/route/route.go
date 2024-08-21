package route

import (
	http "edot-monorepo/services/user-service/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}
func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/login", c.UserController.Login)
	c.App.Post("/api/register", c.UserController.Register)

}

func (c *RouteConfig) SetupAuthRoute() {
	// c.App.Use(c.AuthMiddleware)
}
