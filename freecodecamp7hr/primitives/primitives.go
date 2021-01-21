package main

import "fmt"

func main() {
	fmt.Println("Primitives: Section 2")

	// Boolean type
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	// The 'zero' value for boolean is false
	var zeroVal bool
	fmt.Printf("%v, %T\n", zeroVal, zeroVal)

	var a int = 10 // 1010
	// var b int32 = 12 // Mismatch type, error
	var b int = 3 // 0011
	fmt.Println(a + b)

	// Bitwise operators
	fmt.Println(a & b)  // 0010 = 2
	fmt.Println(a | b)  // 1011 = 11
	fmt.Println(a ^ b)  // XOR (exclusive OR) 1001
	fmt.Println(a &^ b) // and not (you get 1 if both bits are 0)

	x := 8
	fmt.Println(x << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(x >> 3) // 2^3 / 2^3 = 2^0 = 1

	// For floats there is no % operator, nor bitwise or bit shifting operators
	// These operators only work for int types

	// Complex numbers: complex64 (float32 + float32) and complex128 (float64 + float64) (real and imaginary parts)
	var myComplex complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", myComplex, myComplex)
	var imaginaryOnly complex64 = 3i
	fmt.Printf("%v, %T\n", imaginaryOnly, imaginaryOnly)
	// Extract real and imaginary parts with built-in functions
	fmt.Printf("%v, %T\n", real(myComplex), real(myComplex))
	fmt.Printf("%v, %T\n", imag(myComplex), imag(myComplex))
	// complex() functions to create a complex number
	var buildComplex complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", buildComplex, buildComplex)

	// Strings
	s := "Hello world"
	// Treat the string like an array
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])         // prints the ascii code
	fmt.Printf("%v, %T\n", string(s[2]), s[2]) // convert byte to string
	// convert string to a slice of bytes
	message := []byte(s)
	fmt.Printf("%v, %T\n", message, message)

	// rune is a type alias for int32, good for conversions between utf-8 and others
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // 97, int32

}
