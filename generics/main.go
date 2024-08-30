package main

import "fmt"

func add(a, b any) any {
	intA, aIsInt:= a.(int)
	intB, bIsInt := b.(int)

	if	aIsInt && bIsInt {
		return intA + intB
	} 
	return nil
}

func addWithGenerics[T int|float64|string](a, b T) T {
	return a + b
}


func main () {
	// result := add(3,4)
	result := addWithGenerics(3,4)
	result  += 5

	fmt.Println(result)
}