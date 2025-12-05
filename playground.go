package lists

import "fmt"

func main() {

	prices := []float64{10.5, 20.0}

	fmt.Println(prices)

	prices = append(prices, 5.3)

	prices = prices[0:]

	fmt.Println(prices)

	discountedPrices := []float64{12, 13.5, 14.2}

	prices = append(prices, discountedPrices...)

	fmt.Println(prices)

}
