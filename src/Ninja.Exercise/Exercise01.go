package main

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
