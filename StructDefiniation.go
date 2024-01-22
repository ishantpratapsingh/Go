// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type person struct {
	FirstName         string // Declaring String as struct
	LastName          string
	FavouriteIcecream string
	Age               int // Declaring int as struct
}

func main() {
	var a person // Declaring variable as struct without adding values
	fmt.Println(a)
	a = person{"Nihal", "Kumar", "Vanilla", 23}          //using person struct to print values // adding values to predefined struct
	a2 := person{"Bobby", "Bongchie", "Hashbrownie", 26} // Using a new var for struct
	fmt.Println("First Person:", a)
	fmt.Println("Second Person:", a2)
}
