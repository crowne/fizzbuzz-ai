package main

import (
	"fmt"
	"os"
	"strconv"
)

// Add adds two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract subtracts the second number from the first
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply multiplies two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide divides the first number by the second
// Returns an error if the second number is zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <number1> <operation> <number2>")
		fmt.Println("Operations: +, -, *, /")
		os.Exit(1)
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Error parsing first number: %v\n", err)
		os.Exit(1)
	}

	operation := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("Error parsing second number: %v\n", err)
		os.Exit(1)
	}

	var result float64
	var errResult error

	switch operation {
	case "+":
		result = Add(num1, num2)
	case "-":
		result = Subtract(num1, num2)
	case "*":
		result = Multiply(num1, num2)
	case "/":
		result, errResult = Divide(num1, num2)
		if errResult != nil {
			fmt.Printf("Error: %v\n", errResult)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid operation. Please use +, -, *, or /")
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", result)
}