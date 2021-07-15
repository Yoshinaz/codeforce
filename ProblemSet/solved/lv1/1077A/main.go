package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var a, b, k int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &k)
		m := k / 2
		ans := m*a - m*b
		if k%2 == 1 {
			ans += a
		}
		fmt.Println(ans)
	}
}
