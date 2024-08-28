package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	wordsSet := make(map[string]struct{}, len(words))
	for _, word := range words {
		wordsSet[word] = struct{}{}
	}
	fmt.Println(wordsSet)
}
