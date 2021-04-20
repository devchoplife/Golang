package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

//this is a method
//Methods have references to the objects they are being called on which makes them different from functions
//method to get the age of the students
func (s Student) getAge() int {
	return s.age
}

//method to set the age of students
func (s *Student) setAge(age int) {
	s.age = age
}

//function to get average grade
func (s *Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

//function that gets the maximum grade
func (s *Student) getMaximumGrade() int {
	curMax := 0

	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	s1.getAge() //call the method here\

	//Setting the age of thestudent
	fmt.Println(s1)
	s1.setAge(22)
	fmt.Println(s1)

	average := s1.getAverageGrade()
	fmt.Println(average)

	maxGrade := s1.getMaximumGrade()
	fmt.Println(maxGrade)
}
