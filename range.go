package main

import (
	"fmt"
)

func main() {
	//range is  a keyword that allows us to iterate over items or elements inside arrays or slices
	var a []int = []int{1, 3, 4, 5, 7, 23, 4, 9}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//Using range
	for i, element := range a {
		fmt.Printf("%d: %d\n", i, element)
	}

	//Avoid using i  an getting an error
	for _, element := range a {
		fmt.Printf("%d\n", element)
	}

	//Print out duplicates in the slice once
	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && j > i {
				fmt.Println(element)
			}
		}

	}

	// Another way of doing it
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}
