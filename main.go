package main

import (
	"go-price-calculator/cmdmanager"
	"go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0.7, 0.3, 0.15}

	for _, taxRate := range taxRates {
		// fm := *filemanager.New(
		// 	"prices.txt",
		// 	fmt.Sprintf("./results/results_%.0f.json", taxRate*100),
		// )
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		priceJob.Process()
	}
}
