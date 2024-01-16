package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, Ishant")

	// Generate a random number between 0 and 249
	i := rand.Intn(250)
	fmt.Println("The value generated is \t", i)

	// Check the range of the generated value
	if i >= 0 && i < 100 {
		fmt.Println("The value is between 0 and 100")
	} else if i >= 100 && i < 200 {
		fmt.Println("The value is between 100 and 200")
	} else if i >= 200 && i < 250 {
		fmt.Println("The value is between 200 and 250")
	} else {
		fmt.Println("The value is out of range")
	}
}

