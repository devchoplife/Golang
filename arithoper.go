package main

import (
	"fmt"
)

func main() {
	var num1 int = 9
	var num2 int = 4

	answer := num1 / num2
	remainder := num1 % num2

	fmt.Printf("%d remainder %d", answer, remainder)
}
