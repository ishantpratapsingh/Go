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

	fmt.Println(m)         // This will print the entire map
	m["NewKey"] = 99       //Adding a key and value to an existing map
	fmt.Println("The map after adding element: \t", m)
	delete(m, "Dark")      // This will delete a value and key from the map
	fmt.Println("The map after deleting :\t", m)
	m["NewKey"] = 123      //Chaning value of an existing key
	fmt.Println("The map after changing value of an existing key :\t", m)

	for i, v := range m {
		fmt.Println("The value is :\t", v) //value of current iteration prints the value
		fmt.Println("The key is :\t", i)   //Index of current iteration prints the key
	}
}
