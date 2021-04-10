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

	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			continue
		}
		//fmt.Println("Not Valid") - Here is an example of using continue
		//continue makes the for loop keep running until x gets to a thousand
	}

	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			break //break here prints the nfirst number that satisfies this condition
		}
	}
}
