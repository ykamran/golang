package main

import "fmt"

/**
Hands-on exercise #1
Write a program that prints a number in decimal, binary, and hexa decimal
*/

func main() {
	dec := 1024
	var bin = fmt.Sprintf("%b", dec)
	var oct = fmt.Sprintf("%o", dec)
	var hex = fmt.Sprintf("%#x", dec)

	fmt.Println("Decimal Number: ", dec)
	fmt.Println("Binary Number: ", bin)
	fmt.Println("Octal Number: ", oct)
	fmt.Println("Hexa Decimal: ", hex)
}
