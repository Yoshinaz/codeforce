package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, x, y int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &x, &y)
		min, max := minMax(x, y)
		left := 2*(max-min) - 1
		if left < 0 {
			left = 0
		}
		fmt.Println(min*2 + left)
	}
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
