package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pramek008/first-golang/database"
	"github.com/pramek008/first-golang/handlers"
	)

func setupRoutes(app *fiber.App){
	app.Get("/", handlers.Home)
	
	app.Get("/facts", handlers.ListFacts)
	app.Get("/fact/:id", handlers.ShowFact)
	app.Post("/fact", handlers.CreateFact)
	app.Patch("/fact/:id", handlers.UpdateFact)
	app.Delete("/fact/:id", handlers.DeleteFact)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	// app.Get("/", handlers.Home)
	// app.Post("/fact", handlers.CreateFact)

	setupRoutes(app)
	
	app.Listen(":3000")
}
