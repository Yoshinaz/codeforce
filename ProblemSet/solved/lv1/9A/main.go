package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b)
	m := max(a, b)
	prob := 6 - m + 1
	a, b = faction(prob, 6)
	fmt.Printf("%d/%d\n", a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func faction(a, b int) (int, int) {
	if a%2 == 0 {
		a = a / 2
		b = b / 2
	}
	if a%3 == 0 {
		a = a / 3
		b = b / 3
	}
	return a, b
}
