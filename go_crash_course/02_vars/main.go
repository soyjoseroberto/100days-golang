package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint 64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var keyword
	var name = "Jose"
	var age int32 = 38
	const isCool = true
	// isCool = false // Would not work with const

	// Shorthand
	middleName := "Roberto"

	// Single line, multiple assignment
	paternal, maternal := "Rodriguez", "Lezcano"

	shoeSize := 10.5

	fmt.Println(name, middleName, paternal, maternal, age)
	fmt.Printf("%T\n", shoeSize)

}
