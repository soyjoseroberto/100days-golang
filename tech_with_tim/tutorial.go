package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func switchStatement() {
	answer := 10

	// Use switch statement
	switch answer {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fmt.Println("Multiple statements in case 2")
	case 3, 4, 5:
		fmt.Println("3, 4 or 5: multiple cases separated by comma")
	default:
		fmt.Println("answer var did not match any cases")

	}

	// Alternative declaration with switch
	switch {
	case answer > 5 && answer <= 7:
		fmt.Println("Answer is 6 or 7")
	case answer >= 8:
		fmt.Println("Answer greater or equal to 8")
	default:
		fmt.Println("Default case reached")
	}
}

// Arrays, video # 12
func arraysVideo() {
	fmt.Println("Arrays")
	// Declaration. Arrays are fixed in size
	var arr [5]int // int defaults to 0 value
	arr[0] = 100
	arr[4] = 700
	// Other values remain 0
	fmt.Println(arr)

	// Alternative way to init arrays
	arr2 := [3]int{128, 256, 512}
	fmt.Println(arr2)

	// Get length with len() function
	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Printf("Sum for arr2 is %d\n", sum)

	// Multi-dimensional array
	arr2D := [2][3]int{{1, 2, 3}, {3, 6, 9}}
	fmt.Println(arr2D)
	// Print value 9
	fmt.Println(arr2D[1][2])

}

// Slices, video #13
func slicesVideo() {
	fmt.Println("Slices")
	origArr := [5]int{22, 28, 16, 23, 18}
	// Take a slice of the original array with the : operator, last index is not inclusive
	var mySlice []int = origArr[1:4]
	fmt.Println(mySlice)
	fmt.Printf("Slice length %d\n", len(mySlice))
	fmt.Printf("Slice capacity %d\n", cap(mySlice))
	// Re-slice the slice to capacity
	fmt.Println(mySlice[:cap(mySlice)])

	// Declare and init slice
	var s []int = []int{2, 4, 8, 10}
	// It's actually declaring an array under the hood and
	// you are taking a slice of the whole array
	fmt.Println(cap(s))
	// To add a new element to the slice, use append() function
	extendedSlice := append(s, 12)
	fmt.Println(extendedSlice)

	// Slice init with make() function. Pass type and initial size
	makeSlice := make([]int, 3)
	fmt.Println(makeSlice)

}

// range, video #14
func rangeSection() {
	// init a slice
	var grades []int = []int{93, 95, 98, 88, 85, 99, 89}
	// Print the values with regular for loop
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	// Now use range to do the same thing. Get the index and the element
	// Use anonymous variable _ when not needing the index or the element (for _, element := range grades {})
	for index, element := range grades {
		fmt.Printf("%d: %d\n", index, element)
	}
}

// maps, video #15 for key-value pairs
func mapsVideo() {
	var fruitsMap map[string]int = map[string]int{
		"apples":  5,
		"pears":   4,
		"oranges": 10,
	}
	fmt.Println(fruitsMap)

	// Add an element by adding a new key
	fruitsMap["bananas"] = 7
	// Update an element by using its key
	fruitsMap["apples"] = 2
	// Delete using the delete() function
	delete(fruitsMap, "pears")
	fmt.Println(fruitsMap)
	// Check if a key exists on a map
	val, keyExists := fruitsMap["avocados"]
	fmt.Printf("Avocados value %v and key exists: %t\n", val, keyExists)
	// Of course you can get length
	fmt.Println(len(fruitsMap))

	// Alternative with empty map
	// fruitsMap := make(map[string]int)
}

// functions, video #16
func functionsVideo(x, y int) (sum int, mul int) {
	// parameters with same type can be separated by comma
	// you can label return values and have a single return statement
	defer fmt.Println("This executes after return statement. Good for cleaning up")
	sum = x + y
	mul = x * y
	return
}

// functions (advance) and closures, video #17
func execFunction(passedFunc func(int) int) {
	fmt.Println(passedFunc(6))
}

func closuresVideo() {
	test := func(x int) int {
		return x * x
	}
	execFunction(test)
}

func closureFunc(name string) func() {
	return func() {
		fmt.Println(name)
	}
}

// mutable and immutable data types, video #18
// TL;DR: slices and maps are passed by reference so they are mutable (can change)
// Not the case with array and value types
func changeFirstElement(mySlice []int) {
	mySlice[0] = 100
}

// pointers and dereference operator (* and &), video #19
func sayGoodBye(byebye *string) {
	*byebye = "Good bye!!"
}

// Point struct for structs, video #20
type Point struct {
	x int32
	y int32
}

// Circle contains a reference to the Point struct
type Circle struct {
	radius float64
	center *Point
}

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

	// Casting between number types
	var number1 int = 9
	var number2 float64 = 4.0
	// must cast to the same type
	fmt.Printf("Casting to int: %d\n", number1/int(number2))
	// Casting to float
	fmt.Printf("Casting to float: %g\n", float64(number1)/number2)
	// Dividing by zero throws a runtime error, not compile time
	// zeroDivisor := 0
	// fmt.Printf("%d", number1/zeroDivisor)

	// BEDMAS: order of operations (Brackets, Exponent, Division, Multiplication, Addition, Subtraction)

	// Infinite loop
	// for true {
	// }

	// continue keyword; to go back to the top of the loop
	// break keyword; to finish executing the loop

	// switch, video #11
	switchStatement()

	// arrays, video #12
	arraysVideo()

	// slices, video #13
	slicesVideo()

	// range, video #14
	rangeSection()

	// maps, video #15
	mapsVideo()

	// functions, video #16
	sum1, mul1 := functionsVideo(9, 5)
	fmt.Println(sum1, mul1)

	// closures, video #17
	closuresVideo()

	closureFunc("Roberto")()

	// mutable and immutable data types, video #18
	someNumbers := []int{99, 101, 102}
	fmt.Println(someNumbers)
	changeFirstElement(someNumbers)
	fmt.Println(someNumbers)

	// pointers and dereference, video #19
	greeting := "Hello"
	fmt.Println(greeting)
	// Pass the pointer (memory address) where "Hello" is located
	sayGoodBye(&greeting)
	fmt.Println(greeting)

	// structs, video #20
	// When you have an embedded struct (Point) use the pointer to it, not the struct itself
	circle := Circle{2.5, &Point{3, 4}}
	fmt.Println(circle)

}
