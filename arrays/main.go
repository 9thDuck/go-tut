package main

import (
	"errors"
	"fmt"
)

type Product struct {
	title string
	id string
	price float64
}

// func main() {
// 	prices := []float64{1,2,3,4,34,5,6}
// 	prices = append(prices, 1)
// 	fmt.Println(prices)

// 	// featuredPrices := prices[:2]

// 	restOfThePrices := prices[3:5]

// 	// slices are passed by reference
// 	sliceOfTheSlice := restOfThePrices[:1]
// 	sliceOfTheSlice[0] = 5

// 	sliceOfTheSlice = sliceOfTheSlice[:9]



// 	fmt.Println(sliceOfTheSlice)
// }

type product struct {
	title string
	id string
	tags []string
}

func makeProduct (title string, id string, tags []string) (*product, error ) {
 if title == "" || id == "" {
	return &product{}, errors.New("title and id are required")
 }

 return &product{title, id, tags}, nil
} 

func main () {
	// // dyanamic length slice
	// prices := []float64 {12,23,}

	// //This will throw an error while compiling.
	// // prices[15] = 5;

	// // Creates new array behind the scenes, appends to it 1 and return it. 
	// newPrices := append(prices, 1)
	// fmt.Println(newPrices, prices)
	newProduct, err := (makeProduct("asdf", "asdf", []string{"asdf"}))

	if err != nil {
		panic(err)
	}
	newProductVal := *newProduct
	newProductVal.tags = append(newProductVal.tags, "jkl;")

	fmt.Println(newProductVal)

}