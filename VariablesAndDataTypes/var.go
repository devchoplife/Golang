package main

import "fmt"

func main() {
	var name string //implicit varable declaration
	number := 500   //explicit varaiable declaration
	//notice that adding the colon indicates you remove the var keyword

	name = "Michael"
	number = number + 345
	fmt.Println("Hello World, My name is", name, "with the number", number)
	fmt.Printf("%T", number)
}

//int
//uint
//float64
//string
//boolean
//complex
