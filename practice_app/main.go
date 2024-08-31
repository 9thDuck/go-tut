package main

import (
	utils "example.com/practice_app/read_write_to_fs"
	"example.com/practice_app/tax"
)

const outputFilename = "output.json"
const pricesFilename = "prices.txt"

func main() {
	pricesSlice, err := utils.ReadFloat64DataFile(pricesFilename)

	if err != nil {
		panic(err)
	}

	pricesTaxesMap := tax.GetTaxesMapForPriceList(pricesSlice)

	err = utils.WriteJsonToFile(pricesTaxesMap, outputFilename)

	if err != nil {
		panic(err)
	}
}
