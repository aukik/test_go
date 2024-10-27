package server

import (
	"github.com/gofiber/fiber/v2"

	"test_go/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "test_go",
			AppName:      "test_go",
		}),

		db: database.New(),
	}

	return server
}
