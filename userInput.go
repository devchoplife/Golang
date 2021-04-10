package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type theyear you were born: ")

	scanner.Scan() //for getting user input
	//this interprets the user input as a string

	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //the underscore after iput issa variable that captures errors
	//the arguments of the ParseInt refer to thebase and size respectively
	fmt.Printf("You will be %d years old at the end of 2030", 2030-input)
}
