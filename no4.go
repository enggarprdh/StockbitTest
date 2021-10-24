package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	anagram(input)
}

func anagram(data []string) {
	list := make(map[string][]string)
	for _, word := range data {
		index := sortStr(word)
		list[index] = append(list[index], word)
	}
	for _, word := range list {
		for _, w := range word {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func sortStr(word string) string {
	splitWord := strings.Split(word, "")
	sort.Strings(splitWord)
	return strings.Join(splitWord, "")
}
