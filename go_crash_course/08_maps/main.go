package main

import "fmt"

func main() {
	// Define map with make() function passing the types for key and value
	emails := make(map[string]string)

	// Assign key value pairs
	emails["JR"] = "jr@test.com"
	emails["Faye"] = "faye@test.com"
	fmt.Println(emails)
	// Get the length
	fmt.Println(len(emails))
	// Get specific value with key
	fmt.Println(emails["Faye"])

	// Delete from map with delete() function, pass in the key for deletion
	delete(emails, "JR")
	fmt.Println(emails)

	// Declare and define shorthand. Remember comma on last kv pair
	trafficLight := map[string]string{
		"green":  "go",
		"yellow": "slow down",
		"red":    "stop",
	}
	fmt.Println(trafficLight)
}
