package main

import "fmt"

// Entry point of the program
func main() {
	fmt.Println("We are learning functions in Go!")

	// Calling a simple function
	SimpleFunction()

	// Performing arithmetic operations
	sum := add(14, 13)
	fmt.Println("Sum is:", sum)

	product := multiply(11, 11)
	fmt.Println("Total value is:", product)

	quotient := divide(15, 5)
	fmt.Println("The output is:", quotient)

	difference := subtract(15, 5)
	fmt.Println("The difference is:", difference)
}

// SimpleFunction prints a message
func SimpleFunction() {
	fmt.Println("This is a simple function")
}

// add returns the sum of two integers
func add(a, b int) int {
	return a + b
}

// subtract returns the difference of two integers
func subtract(a, b int) int {
	return a - b
}

// multiply returns the product of two integers
func multiply(a, b int) int {
	return a * b
}

// divide returns the quotient of two integers
func divide(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}