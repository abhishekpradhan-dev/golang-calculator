package main

import (
	"fmt"
	"math"
)

func main() {
	var choice int
	var num1, num2 float64
	var operator string

	fmt.Println("Calculator Home")
	fmt.Println("Choose an option:")
	fmt.Println("1. Basic Arithmetic (+, -, *, /)")
	fmt.Println("2. Trigonometric Functions (sin, cos, tan)")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Enter the first number:")
		fmt.Scanln(&num1)

		fmt.Println("Enter an operator (+, -, *, /):")
		fmt.Scanln(&operator)

		fmt.Println("Enter the second number:")
		fmt.Scanln(&num2)

		switch operator {
		case "+":
			fmt.Printf("Result: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("Result: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("Result: %.2f\n", num1*num2)
		case "/":
			if num2 != 0 {
				fmt.Printf("Result: %.2f\n", num1/num2)
			} else {
				fmt.Println("Division by zero is not allowed.")
			}
		default:
			fmt.Println("Invalid operator. Please use one of +, -, *, /.")
		}

	case 2:
		fmt.Println("Enter angle):")
		fmt.Scanln(&num1)

		fmt.Println("Choose a function: sin, cos, tan")
		fmt.Scanln(&operator)

		radians := num1 * math.Pi / 180

		switch operator {
		case "sin":
			fmt.Printf("sin(%.2f) = %.4f\n", num1, math.Sin(radians))
		case "cos":
			fmt.Printf("cos(%.2f) = %.4f\n", num1, math.Cos(radians))
		case "tan":
			if math.Cos(radians) != 0 {
				fmt.Printf("tan(%.2f) = %.4f\n", num1, math.Tan(radians))
			} else {
				fmt.Println("cos must be nonzero")
			}
		default:
			fmt.Println("Invalid function. Please use sin, cos, or tan.")
		}

	default:
		fmt.Println("Invalid choice. Please restart and choose a valid option.")
	}
}
