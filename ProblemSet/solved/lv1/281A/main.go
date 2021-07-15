package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println(strings.ToUpper(input[:1]) + input[1:])
}
