package main

import (
	"fmt"
)

//maps are used for storing key value pairs
func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	} //maps does not recognise the order of the data

	fmt.Println(mp["apple"])
	mp["apple"] = 900
	mp["Tim"] = 800

	delete(mp, "Tim")

	val, ok := mp["apple"] //This says if the key value exists, store the value in val and if it does not make val whatever
	//the default type of the map is i.e if it is int default type will be 0, ok variable is a boolean, It will be set to
	//true if the value exists and false if it does not

	fmt.Println(val, ok)

	//mp2 := make(map[string]int)
	//Typically arrays are slower than maps because they are stored in order
	//Only use arrays when you care about the order of the values

	fmt.Println(len(mp))
}
