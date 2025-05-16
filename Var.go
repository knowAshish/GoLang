package main

import "fmt"

// @title Basic Go Variable Demonstration
// @author Ashish
// @notice This program demonstrates variable declarations, initialization, and printing in Go.
// @dev It includes constants, implicit and explicit declarations, and variable visibility.

func main() {
	// @notice Declaring and initializing variables with different data types
	var name string = "Ashish"
	var age int = 26
	var height float64 = 5.6
	var isAdult bool = true

	fmt.Println(name)    // Prints "Ashish"
	fmt.Println(age)     // Prints 26
	fmt.Println(height)  // Prints 5.6
	fmt.Println(isAdult) // Prints true

	// @notice Declaring variables without initial values (will be zero-valued)
	var a string
	var b int
	var c bool

	fmt.Println(a) // Prints ""
	fmt.Println(b) // Prints 0
	fmt.Println(c) // Prints false

	// @dev Constant declaration
	const _height = 5.6
	fmt.Println(_height)

	// @dev Implicitly typed variable using short declaration
	person := 123
	fmt.Println(person) // Prints 123

	// @dev Demonstrating visibility based on naming convention
	var Public string = "In this PublicVariable is Visibile and can be access from other"
	var private string = "In this privateVariable is only Visibile within the same package"

	fmt.Println(Public)
	fmt.Println(private)
}
