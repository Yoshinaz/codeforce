package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, a, la int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &la)
		count := 0
		for j := 1; j < n; j++ {
			fmt.Fscan(in, &a)
			count += addDense(la, a)
			la = a
		}
		fmt.Println(count)
	}
}

func addDense(a, b int) int {
	min, max := sort2(a, b)
	count := 0
	for min*2 < max {
		count++
		min *= 2
	}
	return count
}

func sort2(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
