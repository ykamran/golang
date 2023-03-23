package main

import "fmt"

/**
Hands-on exercise #2
Using the following operators, write expressions and assign their values to variables:
g. ==
h. <=
i. >=
j. !=
k. <
l. >
Now print each of the variables.
*/

func main() {
	var result = false
	num1 := 1024
	num2 := 2048

	result = num1 == num2
	fmt.Println("num1 == num2: ", result)

	result = num1 <= num2
	fmt.Println("num1 <= num2: ", result)

	result = num1 >= num2
	fmt.Println("num1 >= num2: ", result)

	result = num1 != num2
	fmt.Println("num1 != num2: ", result)

	result = num1 < num2
	fmt.Println("num1 < num2: ", result)

	result = num1 > num2
	fmt.Println("num1 > num2: ", result)
}
