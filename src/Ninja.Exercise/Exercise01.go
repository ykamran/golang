package main

/**
Hands-on exercise #1
	1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
	IDENTIFIERS “x” and “y” and “z”
	a. 42
	b. “James Bond”
	c. true
	2. Now print the values stored in those variables using
	a. a single print statement
	b. multiple print statements
	code: here’s the solution: https://play.golang.org/p/yYXnWXIQNa
*/

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("x: ", x, ", y: ", y, ", z: ", z)

	fmt.Printf("x: %d", x)
	fmt.Printf(", y: %s", y)
	fmt.Printf(", z: %t", z)
}
