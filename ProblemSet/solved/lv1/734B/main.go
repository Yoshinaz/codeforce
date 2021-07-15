package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k2, k3, k5, k6 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &k2, &k3, &k5, &k6)
	i256 := min(k2, k5)
	i256 = min(i256, k6)
	k2 -= i256
	i32 := min(k2, k3)
	fmt.Println(256*i256 + 32*i32)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
