package server

import (
	"fmt"
	"test_go/internal/types"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Get("/health", s.healthHandler)
	s.App.Post("/users", s.createUserHandler) // Added new route
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello",
	}
	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) createUserHandler(c *fiber.Ctx) error {


	var req types.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid request body",
			})
	}
	fmt.Println(req.Username)
	// Validate username
	if req.Username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Username is required",
			})
	}

	// Use the SQLC querier to create the user
	user, err := s.db.Querier().CreateUser(c.Context(), req.Username)
	fmt.Println(err)
	if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to create user",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
