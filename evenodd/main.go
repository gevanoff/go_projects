package main

import (
	"fmt"
	"math"
)

func main() {
	//Populate array
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var nums []int
	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}

	for _, x := range nums {
		if math.Mod(float64(x), 2) == 0 {
			fmt.Println("Number ", x, " is even")
		} else {
			fmt.Println("Number ", x, " is odd")
		}
	}
}
