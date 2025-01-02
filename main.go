package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("calculator-home")
	})

	app.Get("/calculator", func(c *fiber.Ctx) error {
		num1Str := c.Query("x")
		num2Str := c.Query("y")
		operator := c.Query("operator")

		// switch operator {
		// case addition:
		// 	z = x + y
		// case substraction:
		// 	z = x - y
		// case multiplication:
		// 	z = x * y
		// case division:
		// 	z = x / y
		// }

		if num1Str == "" || num2Str == "" || operator == "" {
			return c.Status(400).SendString("Please provide x, y")
		}

		x, err1 := strconv.ParseFloat(num1Str, 64)
		y, err2 := strconv.ParseFloat(num2Str, 64)

		if err1 != nil && err2 != nil {
			return c.Status(400).SendString("x and y should be numbers")
		}

		var z float64

		switch operator {
		case "addition":
			z = x + y
		case "substraction":
			z = x - y
		case "multiplication":
			z = x * y
		case "division":
			z = x / y

		}

		return c.JSON(fiber.Map{
			"answer": z,
		})
	})

	// app.Get("/substraction", func(c *fiber.Ctx) error {
	// 	num1Str := c.Query("x")
	// 	num2Str := c.Query("y")

	// 	if num1Str == "" || num2Str == "" {
	// 		return c.Status(400).SendString("Please provide x, y")
	// 	}

	// 	x, err1 := strconv.ParseFloat(num1Str, 64)
	// 	y, err2 := strconv.ParseFloat(num2Str, 64)

	// 	if err1 != nil || err2 != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"error": "x and y should be numbers.",
	// 		})
	// 	}

	// 	z2 := x - y
	// 	return c.JSON(fiber.Map{
	// 		"x-y": z2,
	// 	})
	// })

	// app.Get("/multiplication", func(c *fiber.Ctx) error {
	// 	num1Str := c.Query("x")
	// 	num2Str := c.Query("y")

	// 	if num1Str == "" || num2Str == "" {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"error": "Please provide x, y .",
	// 		})
	// 	}

	// 	x, err1 := strconv.ParseFloat(num1Str, 64)
	// 	y, err2 := strconv.ParseFloat(num2Str, 64)

	// 	if err1 != nil && err2 != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"error": "x and y should be numbers.",
	// 		})
	// 	}

	// 	z3 := x * y
	// 	return c.JSON(fiber.Map{
	// 		"x*y": z3,
	// 	})
	// })

	// app.Get("/division", func(c *fiber.Ctx) error {
	// 	num1Str := c.Query("x")
	// 	num2Str := c.Query("y")

	// 	if num1Str == "" || num2Str == "" {
	// 		return c.Status(400).SendString("Please provide x, y")
	// 	}

	// 	x, err1 := strconv.ParseFloat(num1Str, 64)
	// 	y, err2 := strconv.ParseFloat(num2Str, 64)

	// 	if err1 != nil && err2 != nil {
	// 		return c.Status(400).JSON(fiber.Map{
	// 			"error": "x and y should be numbers.",
	// 		})
	// 	}

	// 	if y == 0 {
	// 		return c.SendString("y must be non zero")
	// 	}

	// 	z4 := x / y
	// 	return c.JSON(fiber.Map{
	// 		"x/y": z4,
	// 	})

	// })

	log.Fatal(app.Listen(":3000"))
}
