package main

import "fmt"

func main() {

	// Declare a string variable 'name' and initialize it with the value "Ashish"
	name := "Ashish"

	// Declare an integer variable 'age' and initialize it with the value 26
	age := 26

	// Declare a floating-point variable 'height' (type float64 by default)
	// and initialize it with the value 5.62345
	height := 5.62345

	// Print all variables in a simple space-separated format using fmt.Println
	// This prints: Name Ashish Age 26 Height 5.62345
	fmt.Println("Name", name, "Age", age, "Height", height)

	// The following lines are commented out examples showing how to print each
	// variable individually using fmt.Printf with formatting verbs:
	// %s for string, %d for integer, and %.2f for floating-point with 2 decimals.

	// fmt.Printf("Name is %s\n", name)       // Prints: Name is Ashish
	// fmt.Printf("Age is %d\n", age)          // Prints: Age is 26
	// fmt.Printf("Height is %.2f\n", height)  // Prints: Height is 5.62 (rounded to 2 decimals)

	// Print all variables in one formatted string with labels.
	// %s formats the string, %d formats the integer, and %.2f formats the float to 2 decimal places.
	// Output example: Name : Ashish, Age : 26, Height : 5.62
	fmt.Printf("Name : %s, Age : %d, Height : %.2f\n", name, age, height)
}
