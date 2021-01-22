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

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func arraysSection() {
	// Use ... to declare the size of the array
	grades := [...]int{98, 89, 92, 94}
	fmt.Printf("Grades %v\n", grades)

	// slices
	aSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bSlice := aSlice[:]   // slice of all elements
	cSlice := aSlice[3:]  // slice from 4th element (index 3) to the end
	dSlice := aSlice[:6]  // slice of the first 6 elements (not index 6)
	eSlice := aSlice[3:6] // slice of 4th, 5th, and 6th elements

	fmt.Println(aSlice)
	fmt.Println(bSlice)
	fmt.Println(cSlice)
	fmt.Println(dSlice)
	fmt.Println(eSlice)

}

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist) // returns false

	// Store all roles into a single byte (efficient)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin %v\n", isAdmin&roles == isAdmin) // bitwise AND
	fmt.Printf("Is HQ %v\n", isHeadquarters&roles == isHeadquarters)

	arraysSection()

}
