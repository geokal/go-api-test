package server

import (
	"github.com/gofiber/fiber/v2"

	"go-api-test/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-api-test",
			AppName:      "go-api-test",
		}),

		db: database.New(),
	}

	return server
}
