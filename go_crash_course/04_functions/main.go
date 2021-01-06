package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// Argument type applies to both params: int
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Dominic"))
	fmt.Println(getSum(12, 13))
}
