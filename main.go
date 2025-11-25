package main

import (
	"log"
	"project-management/config"
	"project-management/controllers"
	"project-management/database/seed"
	"project-management/repositories"
	"project-management/routes"
	"project-management/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()

	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server running in port: ", port)

	log.Fatal(app.Listen(":" + port))
}
