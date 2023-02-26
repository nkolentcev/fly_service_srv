package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nkolentcev/fly_services_srv/config"
	"github.com/nkolentcev/fly_services_srv/internal/controllers"
)

func init() {
	cfg, err := config.LoadConfig("./../")
	if err != nil {
		log.Fatalln("unable to load app.env")
		log.Fatalln(err.Error())
	}
	config.ConnectDB(&cfg)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/users/:userPin", func(router fiber.Router) {
		router.Get("", controllers.FindUserByPin)
	})

	log.Fatal(app.Listen(":8000"))
}
