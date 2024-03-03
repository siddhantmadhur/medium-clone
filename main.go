package main

import "github.com/gofiber/fiber/v2"



func main() {
    app := fiber.New()

    app.Get("/healthcheck", func(c *fiber.Ctx) error {
        return c.Status(200).SendString("Healt Check: OK!")
    })

    app.Listen(":8080")
}
