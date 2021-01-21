package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

// Enumarated constants (like enum)

const (
	errorSpecialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist) // returns false

}
