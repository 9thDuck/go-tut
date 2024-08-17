package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print(`Enter investment amount: `)
	fmt.Scan(&investmentAmount)

	fmt.Print(`Enter expected return rate: `)
	fmt.Scan(&expectedReturnRate)

	fmt.Print(`Enter tenure in years: `)
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	// fmt.Println("Future value: ", futureValue)
	// fmt.Println("Future real value: ", futureRealValue)

	fmt.Printf("\nOutput\nFuture value: %.2f\nFuture real value: %.2f\n", futureValue, futureRealValue)
}
