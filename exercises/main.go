package main

import "fmt"

func makeItBlue() {
	color := "green"
	// Change the value of color to "blue"
	color = "blue"
	fmt.Println(color)
}

func varToVar() {
	// Declare a variable named color and assign it the value "green"
	color := "green"
	// concatenates the string "dark " to the value of color
	color = "dark " + color
	fmt.Println(color)
}

func multiplyByTwo() {
	// Declare a variable named n and assign it the value of 0(float64)
	n := 0.
	// Assign to n the value of 3.14 multiplied by 2
	n = 3.14 * 2
	// Print the value of n
	fmt.Println(n)
}

func findPerimeter() {
	// Declare variables for perimeter, height, and width
	// assign values 5 to width and 6 to height
	var (
		perimeter int

		width, height int = 5, 6
	)
	// Calculate the perimeter of a rectangle
	perimeter = 2 * (width + height)
	// Print the value of perimeter
	fmt.Println("Perimeter:", perimeter)
}

func swapValues() {
	//Declare variables red and blue and assign them the values "red" and "blue"
	red, blue := "red", "blue"
	// Swap the values of red and blue
	red, blue = blue, red
	// Print the values of red and blue
	fmt.Println("Red:", red, "Blue:", blue)
}

func main() {
	// calls the function makeItBlue, 1st exercise
	// which changes the value of color to "blue"
	makeItBlue()

	// calls the function varToVar, 2nd exercise
	// which changes the value of color to "dark green"
	varToVar()

	// calls the function multiplyByTwo, 3rd exercise
	// which changes the value of n to 6.28
	multiplyByTwo()

	// calls the function findPerimeter, 4th exercise
	// which calculates the perimeter of a rectangle with width 5 and height 6
	findPerimeter()

	// calls the function swapValues, 5th exercise
	// which swaps the values of red and blue
	swapValues()
}
