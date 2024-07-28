package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue,futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)

	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
	// fmt.Printf("Future Value: %v\nFuture Real Value: %v\n", futureValue, futureRealValue)
	// fmt.Printf("Future Value: %.0f\nFuture Real Value: %.0f\n", futureValue, futureRealValue)
	fmt.Print(formattedFV,formattedRFV)

	// fmt.Println("test pointer",&investmentAmount)
	// fmt.Println("test value",investmentAmount)
	// investmentAmount =2000
	// fmt.Println("test value",investmentAmount)
}


func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return 
}