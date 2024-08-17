package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	getInputData(`Enter investment amount: `, &investmentAmount)

	getInputData(`Enter expected return rate: `, &expectedReturnRate)

	getInputData(`Enter tenure in years: `, &years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)

	fmt.Printf("\nOutput\nFuture value: %.2f\nFuture real value: %.2f\n", futureValue, futureRealValue)
}

func getInputData(label string, variable *float64) {
	fmt.Print(label)
	fmt.Scan(variable)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
