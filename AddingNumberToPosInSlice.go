package main

import "fmt"

func main() {
	sl1 := []int{43, 46, 45, 47, 42, 41, 40, 49, 44}
	sl2 := []int{}
	fmt.Println(sl1)

	num := 50
	pos := 6
	for i := 0; i < len(sl1); i++ {

		if i == pos {
			sl2 = append(sl2, num)
		} else {
			sl2 = append(sl2, sl1[i])
		}
	}

	fmt.Println(sl2)
	fmt.Println(num)
}
