package main

import (
	"fmt"
)

func main() {
	x := 3
	for x < 5 {
		fmt.Println(x)
		x++
	}

	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	} //does the same thing as the first for loop

	//break - exits and goes  to te end of the loop
	//continue - skips everything below it and goes back to the begining of the loop

	for {
	}
}
