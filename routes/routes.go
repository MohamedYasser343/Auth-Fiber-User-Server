package routes

import (
	"github.com/MohamedYasser343/handlers"
	"github.com/MohamedYasser343/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Post("/login", handlers.Login)

	r.Get("/user", middleware.Auth, handlers.UserRead)
	r.Get("/user/:id", handlers.USerReadById)
	r.Post("/user", handlers.UserCreate)
	r.Put("/user/:id", handlers.UserUpdate)
	r.Delete("user/:id", handlers.UserDelete)
}
