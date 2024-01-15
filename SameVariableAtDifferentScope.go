package main

import "fmt"

const abc = 10

func anotherone() {
	abc := 20
	fmt.Println("The variable at function level", abc)
}

func main() {
	//abc := 11 - If you uncomment this, the value of abc will be printed as 11, else it will be printed as 10
	fmt.Println("The variable at main level", abc)
	fmt.Println("The variable at function level", abc)

	fmt.Println("The variable at module level", abc)
}
