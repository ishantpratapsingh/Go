package main

import "fmt"

// Person struct embeds the Name struct
type Person struct {
    Name
    Age int
}

// Name struct represents first and last names
type Name struct {
    FirstName string
    LastName  string
}

func main() {
    // Create instances of the Person struct
    personWithMake := Person{
        Name: Name{FirstName: "John", LastName: "Doe"},
        Age:  30,
    }

    personWithoutMake := Person{
        Name: Name{FirstName: "Jane", LastName: "Smith"},
        Age:  25,
    }

    // Access and print values
    fmt.Println("Person with make - First Name:", personWithMake.FirstName)
    fmt.Println("Person with make - Last Name:", personWithMake.LastName)
    fmt.Println("Person with make - Age:", personWithMake.Age)

    fmt.Println("Person without make - First Name:", personWithoutMake.FirstName)
    fmt.Println("Person without make - Last Name:", personWithoutMake.LastName)
    fmt.Println("Person without make - Age:", personWithoutMake.Age)
}
