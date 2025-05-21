package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Print a prompt asking the user for their name
	fmt.Println("Hey What is your name")

	// --- Option 1 (commented out): Using fmt.Scan to read input ---
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	// --- Option 2: Using bufio.NewReader to read full input including spaces ---
	// Create a new buffered reader to read input from standard input (keyboard)
	reader := bufio.NewReader(os.Stdin)

	// Read the input until the newline character ('\n') is encountered
	name, _ := reader.ReadString('\n')

	// Print a greeting message using the input name
	fmt.Println("Hello, Mr.", name)
}
