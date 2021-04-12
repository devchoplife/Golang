package main

import (
	"fmt"
)

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5} //array
	var s []int = x[1:3]                 //slice - does not contain the number of elements

	fmt.Println(len(s))
	fmt.Println(cap(s)) //capacity of the slice
	fmt.Println(s[:cap(s)])

	//making a slice
	var a []int = []int{5, 6, 7, 8, 9} //creates an array first then makes the slice of the whole array
	b := append(a, 10)                 //appends an element to a slice
	fmt.Println(b)

	//Making slices with themake function
	c := make([]int, 5)
	fmt.Println(c)
}
