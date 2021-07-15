package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b, c, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c, &n)
		m := max(a, b)
		m = max(m, c)
		sum := (m - a) + (m - b) + (m - c)

		if n >= sum && (n-sum)%3 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
