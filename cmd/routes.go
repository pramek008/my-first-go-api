package main

import(
	"github.com/gofiber/fiber/v2"
	"github.com/pramek008/first-golang/handlers"
)

// func setupRoutes(app *fiber.App) {
//     app.Get("/", func(c *fiber.Ctx) error {
//         return c.SendString("Div Rhino Trivia App!")
//     })
// }

func setupRoutes(app *fiber.App){
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
}