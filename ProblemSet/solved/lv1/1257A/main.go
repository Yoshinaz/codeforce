package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, x, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &x, &a, &b)
		ans := abs(a - b)
		ans += x
		if ans >= n {
			ans = n - 1
		}
		fmt.Println(ans)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
