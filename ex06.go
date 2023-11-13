package main

import "fmt"

func mergeWordCountMaps(maps []map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for _, wordCount := range maps {
		for word, count := range wordCount {
			mergedMap[word] += count
		}
	}

	return mergedMap
}

func main() {
	wordCountMaps := []map[string]int{
		{"hello": 2, "world": 1},
		{"hello": 1, "goodbye": 1},
		{"world": 3, "foo": 2},
	}

	mergedMap := mergeWordCountMaps(wordCountMaps)

	for word, count := range mergedMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}