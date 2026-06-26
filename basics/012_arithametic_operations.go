package main

import (
	"fmt"
	"math"
)

func main() {
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}