package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/veryshyvelly/task2-backend/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.UserRoutes(app)

	fmt.Println("Listening on port " + port)
	log.Fatalln(app.Listen(":" + port))
}
