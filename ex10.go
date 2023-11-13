package main

import (
	"fmt"
)

func countNumberPairs(numbers []int) map[[2]int]int {
	pairCount := make(map[[2]int]int)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			pairCount[pair]++
		}
	}

	return pairCount
}

func main() {
	numberSlice := []int{1, 2, 3, 1, 2, 4, 3, 1, 2, 3}

	pairFrequency := countNumberPairs(numberSlice)

	for pair, frequency := range pairFrequency {
		fmt.Printf("%v: %d\n", pair, frequency)
	}
}
}