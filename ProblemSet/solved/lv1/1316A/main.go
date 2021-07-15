package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, m, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		sum := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			sum += a
		}
		fmt.Println(min(sum, m))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
