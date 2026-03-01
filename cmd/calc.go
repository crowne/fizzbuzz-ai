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
	if len(os.Args) < 4 {
		fmt.Println("Usage: calc <number1> <operation> <number2>")
		fmt.Println("Operations: +, -, *, /")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	if err1 != nil {
		fmt.Printf("Error parsing first number: %v\n", err1)
		return
	}

	operation := os.Args[2]

	num2, err2 := strconv.ParseFloat(os.Args[3], 64)
	if err2 != nil {
		fmt.Printf("Error parsing second number: %v\n", err2)
		return
	}

	var result float64
	var err error

	switch operation {
	case "+":
		result = Add(num1, num2)
	case "-":
		result = Subtract(num1, num2)
	case "*":
		result = Multiply(num1, num2)
	case "/":
		result, err = Divide(num1, num2)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	default:
		fmt.Println("Invalid operation. Please use +, -, *, or /")
		return
	}

	fmt.Printf("Result: %v\n", result)
}