package main

import "fmt"

func doubleNum(num int) int {
	return num * 2
}

func tripleNum(num int) int {
	return num * 2
}

type TransformCb func(int) int

func getDoubleTransformerFunction() TransformCb {
	return doubleNum
}

func transformNumsSlice(nums *[]int, transformCb TransformCb) {
	numsVal := *nums
	for idx := range numsVal {
		numsVal[idx] = transformCb(numsVal[idx])
	}
}

func main() {
	nums := []int{1, 2, 3, 4}

	transformNumsSlice(&nums, getTransformerFunction())
	fmt.Println(nums)
	transformNumsSlice(&nums, tripleNum)
	fmt.Println(nums)

}
