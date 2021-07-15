package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var p, a, b, c int64
	var t int

	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &p, &a, &b, &c)
		ans := findClass(p, a)
		ans = min(ans, findClass(p, b))
		ans = min(ans, findClass(p, c))
		fmt.Println(ans)
	}
}

func findClass(p, a int64) int64 {
	remind := p % a
	if remind == 0 {
		return 0
	}
	return a - remind
}
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
