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
	m := min(a, b)
	ans := int64(1)
	for i := 2; i <= m; i++ {
		ans *= int64(i)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
