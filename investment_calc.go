package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

// var investmentAmount, years float64 = 1000, 10 // cannot declare with := outside a func, must use const or var
// var expectedReturnRate = 5.5

func main() {
	// const inflationRate = 6.5										declare const and vars outside of fxn to be global and used anywhere, as abover
	var investmentAmount, years float64 = 1000, 10 // := used without var
	expectedReturnRate := 5.5                      // := should be used often when wanting the inferred type to stay, don't need an explicit type

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// vars can be re-assigned, makes sense in other programs where fetch user input: 	investmentAmount = 2000
	// fmt.Print("Investment Amount: ")					using our own defined fxn for print
	outputText("Investment Amount: ")
	// fmt.Scan allows user input, needs to pass pointer (&), does not work with constants
	fmt.Scan(&investmentAmount) //& allows scan to populate

	//prompt and print years
	fmt.Print("Years: ")
	fmt.Scan(&years)

	//prompt and print expected return rate
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// take into account inflation
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

// for a string placed on multiple lines, need backtic instead of double quotes,
// and \n is treated as test, and actual line breaks are treated as line breaks in string text

// build own utility print fxn
func outputText(text string) { // can then use this in place of fmt.Print above
	fmt.Print(text)
}

// ^add return value types here	(if using more than one must wrap in parenthesis (int, int))
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) { // create own fxn math.Pow to return (can return multiple values by separating with comma)
	// fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// rfv := fv / math.Pow(1+inflationRate/100, years)
	fv = investmentAmount + math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv // instead of writing return, can also describe in return value types above (fv float64, rfv float64) thus you do not need :=, can just assign as =

}
