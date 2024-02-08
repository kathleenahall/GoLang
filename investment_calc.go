package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10
	var expectedReturnRate := 5.5														// := should be used often when wanting the inferred type to stay, don't need an explicit type


	var futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
