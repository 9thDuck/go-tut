package main

import "fmt"

// like rest op
func sumInts(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(sumInts(1, 2, 3, 4, 5, 6, 7))
}
