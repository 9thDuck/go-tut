package tax

import (
	"fmt"
)

var taxRates = []float64{
	0, 0.1, 0.2, 0.3,
}

type PriceTaxesMap map[string][]float64

func GetTaxesMapForPriceList(prices []float64) PriceTaxesMap {
	priceTaxesMap := PriceTaxesMap{}

	for _, taxRate := range taxRates {
		pricesForTaxRate := []float64{}

		for _, price := range prices {
			pricesForTaxRate = append(pricesForTaxRate, price*(1+taxRate))
		}

		priceTaxesMap[fmt.Sprint(taxRate)] = pricesForTaxRate
	}

	return priceTaxesMap
}
