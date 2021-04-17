package main

import (
	"fmt"
)

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {
	fmt.Println("Hello, Welcome to Mutable and Immutable Datatypes in Golang")
	//immutable - Not changeable
	//Mutable - changeable

	//immutable example
	//Modification here is individual
	var x int = 5
	y := x
	y = 7
	fmt.Println(x, y)
	//int, string, float, boolean are immutable data types

	//Mutable examples
	//A slice isa mutable datatype
	//If a is modified it affects b
	var a []int = []int{3, 4, 5}
	b := a
	b[0] = 100
	fmt.Println(a, b)

	//maps are also mutable data typesbut they behave weird
	var c map[string]int = map[string]int{"hello": 3}
	d := c
	d["y"] = 100
	fmt.Println(c, d)

	//Arrays
	//Arrays are mutable data types
	var e [2]int = [2]int{3, 4}
	f := e
	f[0] = 100
	fmt.Println(e, f)

	//Slices, Maps and arrays are mutable datatypes because they can change
	//Mutable datatypes change the data anywhere they are being referenced

	var g []int = []int{3, 4, 5}
	fmt.Println(g)
	changeFirst(g)
	fmt.Println(g)
}
