// Take Two maps and create a struct for both the maps.

package main

import "fmt"

type person struct {
	FirstName         string // Declaring String as struct
	LastName          string
	FavouriteIcecream string
	Age               int // Declaring int as struct
}

func main() {
	//Declaring map of struct type with make function is below
	map1 := make(map[string]person)
	map1["Nihal"] = person{FirstName: "Nihal", LastName: "Yadav", Age: 24}

	//Declaring map of struct type without make function is below

	map2 := map[string]person{
		"Nihal": person{FirstName: "Nihal", LastName: "Kumar", Age: 23},
		"Bobby": person{FirstName: "Bobby", LastName: "Bongchie", Age: 25},
	}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println("-------------------")
	fmt.Println(map2["Bobby"].LastName) //Printing a value with specific key of map2
	fmt.Println("---------------")
	juststruct := person{"Ishant", "Singh", "chocolate", 24} //Declaring a struct
	fmt.Println(juststruct)                                  //Printing the struct
}
