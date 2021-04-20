package main

import "fmt"

//struct stands for structure. It is a custom type
//The code below means type followed by name of the type (starting with BLOCk letter) then struct keyword
//We are basically creating anew data type
type Point struct {
	//list fields that will be part of this struct
	x int32
	y int32
}

//Embedded struct
//it is important to use pointers and dereference hen using embedded structs
type Circle struct {
	radius float64
	origin *Point
}

//Using the pointer to change the value of the Point
func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	var p1 Point = Point{1, 2} //this means Point has x = 1 and y = 2
	var p2 Point = Point{-5, 7}

	fmt.Println(p1.y, p2.y)

	//Another way of declaring a point
	//p3 := Point{4, 5}

	//How t declare just one value in a struct
	p4 := Point{x: 5} //This code sets x to 5 and y to the default value of int32
	fmt.Println(p4)

	p5 := &Point{y: 3}
	fmt.Println(p5)
	changeX(p5)
	fmt.Println(p5)

	//When using pointers with structs, you do not need to use the asterisk to reference it
	p6 := &Point{3, 4}
	p6.x = 100
	fmt.Println(p6)

	//embedded struct
	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1)

	fmt.Println(c1.origin.x)
	c1.origin.x = 45
}
