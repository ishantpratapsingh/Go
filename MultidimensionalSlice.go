package main

import "fmt"

func main() {
	slice1 := []int{2, 5, 3, 4, 1, 6}
	slice2 := []int{8, 9, 13, 11, 10, 12, 19, 18}

	slicemd := [][]int{slice1, slice2} // This is a multidimensional Slice
	fmt.Println(slicemd)

	//Above is one way of declaring multidimensional slice and below is another way of doing that

	fmt.Println("-----------")

	var matrix [][]int // Declaring a 2D slice

	matrix = append(matrix, []int{1, 2, 3}) // The slice that was declared earlier is getting values assigned.
	matrix = append(matrix, []int{4, 5, 6})

	fmt.Println(matrix[1][2]) // This will print a number at the position [row][column]
	fmt.Println(matrix)
}
