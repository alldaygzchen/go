package prices

import (
	"fmt"

	"chap9/project/conversation"
	"chap9/project/iomanager"
)

type TaxIncludedPricesType map[string]string

// not an embedded struct
type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices"`
	TaxIncludedPrices TaxIncludedPricesType `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error{
	
	
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversation.StringsToFloats(lines)

	if err != nil {
		return err
	}


	// for lineIndex, line := range lines {
	// 	floatPrice, err := strconv.ParseFloat(line, 64)

	// 	if err != nil {
	// 		fmt.Println("Converting price to float failed.")
	// 		fmt.Println(err)
	// 		file.Close()
	// 		return
	// 	}

	// 	prices[lineIndex] = floatPrice
	// }

	job.InputPrices = prices
	return nil
}



func (job *TaxIncludedPriceJob) Process() error{

	err:=job.LoadData()

	if err != nil {
		return err
	}
	result := make(TaxIncludedPricesType)

	for _, price := range job.InputPrices {
		taxIncludedPrice :=price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)

}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}