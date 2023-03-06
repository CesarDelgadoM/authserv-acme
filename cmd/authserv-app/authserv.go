package main

import (
	"log"

	"github.com/CesarDelgadoM/authserv-acme/database/memory"
	"github.com/CesarDelgadoM/authserv-acme/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Init database users
	db := memory.NewMemoryDB()

	// app Fiber
	app := fiber.New(fiber.Config{
		AppName: "Auth Service ACME v1.0.0",
	})

	// Init authserv handler
	handler.NewAuthHandler(app, db)

	log.Fatal(app.Listen(":8081"))
}
