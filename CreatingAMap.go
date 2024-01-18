// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	m := map[string]int{ //Declaring a map of Key (string) and value (int)
		"Todd":  42,
		"Mac":   55,
		"Dark":  42,
		"Bark":  22,
		"Shark": 11}

	fmt.Println(m) // This will print the entire map

	/*This will not work if you have to print key and value in map. Only for range will work.
	for i := 0; i < len(m); i++ {
		fmt.Println("The key is : ", m[i])
	} */

	for i, v := range m {
		fmt.Println("The value is :\t", v) //value of current iteration prints the value
		fmt.Println("The key is :\t", i)   //Index of current iteration prints the key
	}
}
