package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "selamat malam"
	words := strings.Split(kata, "")
	wordCount := make(map[string]int)
	for _, word := range words {
		fmt.Println(word)
		_, exists := wordCount[word]
		if exists {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}

	fmt.Printf("Hasil = ")
	for key, value := range wordCount {
		fmt.Printf("%s: %d ", key, value)
	}
}
