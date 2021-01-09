package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Declare var with implicit walrus operator
	x := 10

	// Print the type of variable with verb %T
	fmt.Printf("%T\n", x)

	// All types have a default value
	var y int
	var firstName string
	var isHappy bool
	// 0, empty string, false
	fmt.Printf("Default values: %v, %v, %v\n", y, firstName, isHappy)

	z := 15
	fmt.Printf("Binary for 15: %b\n", z)

	// Save output with Sprintf
	out := fmt.Sprintf("Padding is cool %07d", 23)
	fmt.Println(out)

	// Reading input from the console
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type something: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %q\n", input)

	// How old will you be in 2021
	fmt.Printf("Year you were born: ")
	scanner.Scan()
	birthYear, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will turn %d in 2021\n", 2021-birthYear)
}
