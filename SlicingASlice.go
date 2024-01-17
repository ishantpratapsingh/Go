package main

import "fmt"

func main() {

	sl1 := []int{43, 46, 45, 47, 42, 41, 40, 49, 44}

	fmt.Println(sl1[5:])
	fmt.Println(sl1[:4])
	fmt.Println(sl1[0:6])
	fmt.Println(sl1[6:9])
}
