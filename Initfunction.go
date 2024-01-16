package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is a point where main has executed")
}

func init() {
	x := rand.Intn(300)
	switch {
	case x > 0 && x <= 100:
		fmt.Println("The number is between 0 and 100 : \t", x)
	case x >= 100 && x < 200:
		fmt.Println("The number is between 100 and 200 : \t", x)
	case x >= 200 && x < 250:
		fmt.Println("The number is between 200 and 250 : \t", x)
	default:
		fmt.Println("You are dumb if this prints", x)
	}
}
