package main

import (
	"chap10/projects/filemanager"
	"chap10/projects/prices"
	"fmt"
)

/*

....

*/



func main(){
	
	
	// prices := []float64{10.0, 20.0, 30.0}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool,len(taxRates))
	errorChans := make([]chan error,len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index],errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }

	}
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

}