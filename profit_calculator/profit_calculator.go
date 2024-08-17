package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	getInputData(`Enter revenue: `, &revenue)

	getInputData(`Enter expenses: `, &expenses)

	getInputData(`Enter tax rate: `, &taxRate)

	earningsBeforeTax, earningsAfterTax, ratio := calculateProfitsAndRatio(revenue, expenses, taxRate)

	fmt.Printf("\nOutput\nEarnings before tx: %.2f\nEarnings after tax: %.2f\nRatio: %.2f\n", earningsBeforeTax, earningsAfterTax, ratio)
}

func getInputData(label string, variable *float64) {
	fmt.Print(label)
	fmt.Scan(variable)
}

func calculateProfitsAndRatio(revenue, expenses, taxRate float64) (earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax - (earningsBeforeTax * taxRate / 100)
	ratio = earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
