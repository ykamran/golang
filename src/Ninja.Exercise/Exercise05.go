package main

import "fmt"

/**
Hands-on exercise #5
Building on the code from the previous example
	1. At the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
       The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
		a. eg:
			type UserId int
			var x UserId
			var y int
	2. in func main
		a. this should already be done
			i. print out the value of the variable “x”
			ii. print out the type of the variable “x”
			iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
			iv. print out the value of the variable “x”
		b. now do this
			i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
				1. then use the “=” operator to ASSIGN that value to “y”
				2. print out the value stored in “y”
				3. print out the type of “y”
*/

type UserId int

var x UserId
var y int

func main() {
	fmt.Printf("x: %d, and type of x: %T\n", x, x)
	x = 100
	fmt.Printf("x: %d, and type of x: %T\n", x, x)

	y = int(x)
	fmt.Printf("y: %d, and type of y: %T\n", y, y)
}
