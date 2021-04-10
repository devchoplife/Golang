package main

import "fmt"

func main() {
	x := 8

	val := true || false && true
	val2 := (true || false) && !false || x < 9
	val3 := val || false

	fmt.Printf("%t, %t, %t", val, val2, val3)
}
