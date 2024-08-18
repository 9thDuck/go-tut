package main

import (
	"fmt"
)

// input gets an array from the user.
func input() (m int, array []int) {
	fmt.Print("Enter the size of the array, n, and the difference, m: ")
	var n int
	_, err := fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter the array as a space seperated string: ")
	array = make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&array[i])
	}

	return m, array
}

func main() {
	m, array := input()
	fmt.Println(m, array)
}
