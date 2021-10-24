package main

import (
	"fmt"
	"regexp"
)

func main() {
	inputNo3 := "(a(sawew)"
	result := findFirstStringInBracket(inputNo3)
	fmt.Println(result)
}

func findFirstStringInBracket(str string) string {
	re := regexp.MustCompile(`^[\(]*(.*?)[\)]*$`)
	var result = re.ReplaceAllString(str, "${1}")
	return result
}
