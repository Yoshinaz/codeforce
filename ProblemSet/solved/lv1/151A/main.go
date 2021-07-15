package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, l, c, d, p, nl, np int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k, &l, &c, &d, &p, &nl, &np)
	liquid := k * l / nl
	limes := c * d
	salts := p / np
	ans := min(liquid, limes)
	ans = min(ans, salts)
	fmt.Println(ans / n)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
