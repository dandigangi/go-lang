package main

import "fmt"

// Variable declarations w/ initializers w/ factored block
var (
	a, b string
	c    string = "Sam"
	d    bool
)

func main() {
	// Can use block scope to overwrite previously defined vars at package(?) level
	var a int = 5
	var xyz float32 = float32(a)
	var abc uint = uint(xyz)

	// Short hand, implict type - these can't be done outside of a function body
	e := -42
	k := false

	// Types default to 0 values
	fmt.Println(a, b, c, d, e, k, xyz, abc)

	// Implicit type figured out on its own, pick an mno declaration and run
	// mno := 5
	// mno := 5.251
	mno := 5.2513461 + 0.3i

	fmt.Printf("mno type is %T\n", mno) // Incl. string interpolation
}
