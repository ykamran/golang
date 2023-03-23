package main

import "fmt"

/**
Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

const AES_KEY string = "123456"
const USER_ID = 0

const COUNTER = iota

const (
	YEAR2023 = 2023 + iota
	YEAR2024 = YEAR2023 + iota
	YEAR2025 = YEAR2024 + iota
)

func main() {
	fmt.Printf("AES Key: %s and Type %T\n", AES_KEY, AES_KEY)
	fmt.Printf("USER ID:  %d and Type %T\n", USER_ID, USER_ID)

	fmt.Printf("YEAR2023:  %d and Type %T ", YEAR2023)
	fmt.Printf("YEAR2024:  %d and Type %T", YEAR2024)
	fmt.Printf("YEAR2025:  %d and Type %T", YEAR2025)
}
