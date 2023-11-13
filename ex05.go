package main

import "fmt"

func charFrequency(str string) map[rune]int {
	frequencyMap := make(map[rune]int)

	for _, char := range str {
		frequencyMap[char]++
	}

	return frequencyMap
}

func main() {
	text := "hello, world!"

	frequencyMap := charFrequency(text)

	for char, frequency := range frequencyMap {
		fmt.Printf("%c: %d\n", char, frequency)
	}
}