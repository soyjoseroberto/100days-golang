package main

import "fmt"

func main() {
	// Arrays: Fixed size
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Declare and assign shorthand
	colorsArr := [2]string{"red", "blue"}

	fmt.Println(colorsArr)

	// Slices: Dynamic size arrays
	carsArr := []string{"Honda", "Toyota", "Ford", "GMC"}

	fmt.Println(carsArr)

	// Get array length
	fmt.Println(len(carsArr)) //4

	// Get a range of elements, second number not inclusive
	fmt.Println(carsArr[1:3])

}
