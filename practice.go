package main

// Factored block - Can also be applied to var declarations (see var-inits.go)
// import (
// 	"fmt"
// )

import "fmt"

// Add
// Function w/ combined matching types
func add(x, y int) int {
	return x + y
}

// Swap
// Return 2 ints on fx call
func swapNums(x, y int) (int, int) {
	return y, x
}

// Returns
// In short fx's, returns will return named vars
func namedReturnsSubtract(num1 int, num2 int) (a, b, c int) {
	a = num1 - 10
	b = num2 - 5
	c = a - 5
	return
}

func main() {
	// Hello World
	fmt.Println("Hello World")

	// Add
	fmt.Println(add(2, 3))

	// Swap Nums
	// Returns 2 ints, 2 vars must be assigned
	a, b := swapNums(1, 2)
	fmt.Println(a, b)

	// Returns
	fmt.Println(namedReturnsSubtract(40, 20))
}
