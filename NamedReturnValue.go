package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return 
}

func main() {
	fmt.Println(split(17))
}

//This will return 10 and 7 which is same as return x , y so this type of return is called named return
