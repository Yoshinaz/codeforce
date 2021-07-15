package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		min, max := minmax(a, b)
		_, max = minmax(min*2, max)
		fmt.Println(max * max)
	}
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
