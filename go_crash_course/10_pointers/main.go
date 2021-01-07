package main

import "fmt"

func main() {
	a := 5
	b := &a // b has the memory addr of a
	fmt.Println(a, b)

	// print the types
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read the val from a reference variable
	fmt.Println("b equals", *b)
	// This reduntant but shows what you are doing with dereferencing
	fmt.Println("a equals", *&a)

	// Change a through b
	*b = 10
	fmt.Println("Now a equals", a)
}
