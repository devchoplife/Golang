package main

import (
	"fmt"
)

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("You can ride")
	} else if age <= 14 {
		fmt.Println("You cannot ride")
		fmt.Printf("Wait %d years", 18-age)
	} else {
		fmt.Println("You can ride with a parent")
	}
}
