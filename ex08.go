package main

import (
	"fmt"
)

func calculateExpenses(balance map[string]float64) map[string]float64 {
	total := 0.0
	for _, amount := range balance {
		total += amount
	}

	average := total / float64(len(balance))

	adjustedBalance := make(map[string]float64)
	for name, amount := range balance {
		adjustedBalance[name] = amount - average
	}

	return adjustedBalance
}

func main() {
	expenses := map[string]float64{
		"Alice":  20.0,
		"Bob":    15.0,
		"Charlie": 30.0,
		"David":  10.0,
	}

	adjustedBalance := calculateExpenses(expenses)

	for name, amount := range adjustedBalance {
		fmt.Printf("%s: %.2f\n", name, amount)
	}
}