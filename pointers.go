package main

import (
	"fmt"
)

func changeValue(str *string) {
	*str = "Changed!!!"
}

func changevalue2(str string) {
	str = "Changed the Second methos!!!"
}

func main() {
	//There are 2 operators
	// &(ampersand) - Pointer/reference  (memory location of the data )
	// *(asterisk) - dereference

	x := 7
	fmt.Println(&x) //This prints he location of the variable x in the computer's memory
	y := &x
	fmt.Println(x, y)
	*y = 8 //the asterisk simply means access the value in that memory location and change it to x. In this process
	//x iis modified]
	fmt.Println(x, y)

	//This works for any variable(mutable and immutable data types inclusive)

	toChange := "Hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	changevalue2(toChange)
	fmt.Println(toChange) //This method does not change it because it does not point to the pointer

	toChange2 := "Hello!!!"
	var pointer *string = &toChange2
	fmt.Println(*pointer)
	fmt.Println(pointer)
	fmt.Println(pointer, &pointer)
}
