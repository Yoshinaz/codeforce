package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &input)
	input = strings.ToLower(input)
	var out strings.Builder
	for _, v := range input {
		if !isVowel(v) {
			out.WriteString(".")
			out.WriteRune(v)
		}
	}
	fmt.Println(out.String())
}

func isVowel(a rune) bool {
	switch a {
	case 'a', 'o', 'y', 'e', 'u', 'i':
		return true
	default:
		return false
	}
}
