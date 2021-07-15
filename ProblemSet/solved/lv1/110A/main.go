package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	x := strconv.Itoa(count47(input))
	if len(x) == count47(x) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func count47(input string) int {
	count := 0
	for _, v := range input {
		if v == '4' || v == '7' {
			count++
		}
	}
	return count
}
