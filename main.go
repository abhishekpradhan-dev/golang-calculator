package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main(
	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Welcome to calculator")
    })


	app.Get("/calculate", unc(c *fiber.Ctx) error {
		num1Str := c.Query("num1")
		num2Str := c.Query("num2")
		operator := c.Query("operator")
	})

	


    log.Fatal(app.Listen(":3000"))
)