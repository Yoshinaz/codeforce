package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	count := 0
	for _, v := range input {
		if unicode.IsUpper(v) {
			count++
		}
	}
	if count > len(input)-count {
		fmt.Println(strings.ToUpper(input))
	} else {
		fmt.Println(strings.ToLower(input))
	}
}
