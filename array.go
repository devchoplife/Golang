package main

import (
	"fmt"
)

func main() {
	//Golang array indexing starts a t 0
	var arr [5]int

	arr[0] = 100
	arr[4] = 900

	fmt.Println(arr[0])

	//Another way of declaring arrays
	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	sum := 0

	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}

	fmt.Println(sum)

	//Multidimensional arrays
	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D[0][2])
}
