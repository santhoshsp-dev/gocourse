package main

import (
	"fmt"
	"slices"
)

func main() {

	// var number []int
	// var number1 = []int{5,6}
	// number2 := []int{7,8,9}

	// slice := make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}

	slice1 := a[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 5, 6, 7)
	fmt.Println("Slice1: ", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Slice Copy: ", sliceCopy)

	// sliceCopy[1] = 30
	// fmt.Println("Slice Copy1: ", sliceCopy)

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 is equal to slicecopy")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Outer slice at index [i: %d], and in inner slice index of  [j: %d], Adding value [i+j: %d]\n", i, j, i+j)
			// fmt.Println(twoD[i][j])
		}

	}
	fmt.Println(twoD)

	slice2 := slice1[2:4]
	fmt.Println("Slice2: ", slice2)
	fmt.Println("Slice1: ", slice1)
	fmt.Println("The capacity of slice1 is: ", cap(slice1))
	fmt.Println("The capacity of slice2 is: ", cap(slice2))
}
