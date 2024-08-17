package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print(`Enter revenue: `)
	fmt.Scan(&revenue)

	fmt.Print(`Enter expenses: `)
	fmt.Scan(&expenses)

	fmt.Print(`Enter tax rate: `)
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses

	earningsAfterTax := earningsBeforeTax - (earningsBeforeTax * taxRate / 100)

	ratio := earningsBeforeTax / earningsAfterTax

	// fmt.Println(earningsBeforeTax)
	// fmt.Println(earningsAfterTax)
	// fmt.Println(ratio)

	fmt.Printf("\nOutput\nEarnings before tx: %.2f\nEarnings after tax: %.2f\nRatio: %.2f\n", earningsBeforeTax, earningsAfterTax, ratio)

}
