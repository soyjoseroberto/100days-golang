package main

import "fmt"

func main() {
	ids := []int{33, 22, 28, 16, 23, 18}

	for i, id := range ids {
		fmt.Printf("%d ID - %d\n", i, id)
	}

	// Not using the index
	for _, id := range ids {
		fmt.Printf("%d is the value\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	// range with maps
	trafficLights := map[string]string{
		"green":  "go",
		"yellow": "slow down",
		"red":    "stop",
	}

	// loop through key value pair
	for k, v := range trafficLights {
		fmt.Printf("%s: %s\n", k, v)
	}

	// loop through just the key
	for k := range trafficLights {
		fmt.Println("Key is: " + k)
	}
}
