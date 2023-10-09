package main

import (
	"go-api/controller"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App, logger *logrus.Logger, client *http.Client) {
	kodePosController := controller.NewKodePosController(client)

	app.Get("/", status)

	//external api requester
	app.Get("/api/kodepos", kodePosController.GetAllKodePos)
}

func main() {
	app := fiber.New()
	// app.Use(logger.New())
	// Default config
	app.Use(cors.New())
	//initiate logger
	log := logrus.New()
	client := &http.Client{}
	setupRoutes(app, log, client)
	//memanggil func restricted
	app.Listen(":8080")
}
