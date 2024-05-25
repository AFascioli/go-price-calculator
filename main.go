package main

import (
	"fmt"
	"go-price-calculator/filemanager"
	"go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0.7, 0.3, 0.1}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := *filemanager.New(
			"prices.txt",
			fmt.Sprintf("./results/results_%.0f.json", taxRate*100),
		)
		// cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

	}

	for index := range taxRates {
		// will take the first value emmited by any of the channel cases.
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("done!")
		}
	}
}
