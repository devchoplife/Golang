package main

import (
	"fmt"
)

//the int outside is the type to return the value as
func add(x, y int) (int, int) {
	return x + y, x - y
}

//Method 2
func add2(x, y int) (z1 int, z2 int) {
	defer fmt.Println("Hello")
	//The defer function defers whatever is there until the function has finished excuting i.e at the end of the function
	z1 = x + y
	z2 = x - y
	fmt.Println("before return")
	return
}

func main() {
	ans1, ans2 := add2(6, 8)
	fmt.Println(ans1, ans2)
}
