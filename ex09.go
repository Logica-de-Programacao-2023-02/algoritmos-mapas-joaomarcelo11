package main

import (
	"fmt"
)

func generateFibonacci(n int) map[int]int {
	fibonacciMap := make(map[int]int)

	a, b := 0, 1
	index := 0

	for b <= n {
		fibonacciMap[index] = a
		a, b = b, a+b
		index++
	}

	return fibonacciMap
}

func main() {
	n := 50

	fibonacciSequence := generateFibonacci(n)

	for index, number := range fibonacciSequence {
		fmt.Printf("%d: %d\n", index, number)
	}
}