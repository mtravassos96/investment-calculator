package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64
	
	fmt.Print("Your investment amout (USD): ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years of investment: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

