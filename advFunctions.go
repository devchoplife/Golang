package main

import (
	"fmt"
)

func test() {
	fmt.Println("hello")
}

func test3(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func main() {
	x := test
	x()

	//creating a function inside a function
	//main is already a function so we will define another function here
	test2 := func(x int) int {
		return x * -1
	}(8) //this calls the function inline
	fmt.Println(test2)

	test3(test2)
}
