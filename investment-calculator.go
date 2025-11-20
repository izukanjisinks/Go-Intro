package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	var investmentAmount float64
	exptectedRateOfReturn := 5.5
	var years float64 = 10

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+exptectedRateOfReturn/100, years)

	featureRealVale := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)

	fmt.Println(featureRealVale)

}
