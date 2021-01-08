package main

import (
	"fmt"
	"strconv"
)

// Person Define Person struct. Yu can declare same type vars in one line
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver: does not change anything)
func (p Person) greet() string {
	return "Hi, my name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer  receiver: changes something)
func (p *Person) hasBirthday() { // Does not need return value type
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}
}
func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Jose",
		lastName:  "Rodriguez",
		city:      "Chicago",
		gender:    "m",
		age:       37}

	// Alternative init without properties names
	person2 := Person{
		"Faye",
		"McCurdy",
		"Chicago",
		"f",
		37}
	fmt.Println(person1)
	fmt.Println(person2)
	// Access a property with the . operator
	fmt.Println(person1.city)
	// Change a value
	person1.age++
	fmt.Println(person1)

	// Use greeting method (value receiver)
	fmt.Println(person1.greet())

	// Use hasBirthday method (pointer receiver)
	person2.hasBirthday()
	fmt.Println(person2.greet())

	// Use getMarried method (pointer receiver)
	person2.getMarried("Rodriguez")
	fmt.Println(person2.greet())

}
