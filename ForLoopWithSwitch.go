// Write a program to print numbers from 0 to 100 and use switch case to run this loop 100 times


package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println("The number in this iteration is :  ", i)

		switch {

		case true:
			fmt.Println("Number using switch is :  \t ", i)

		}
	}
}



