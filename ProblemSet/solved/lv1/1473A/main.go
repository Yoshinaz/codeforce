package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, d, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &d)
		m1 := 1000
		m2 := 1000
		ans := true
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			m1, m2 = getMin(m1, m2, a)
			if a > d {
				ans = false
			}
		}
		if ans || m1+m2 <= d {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func getMin(a, b, c int) (int, int) {
	if c < a {
		return c, a
	}
	if c < b {
		return a, c
	}
	return a, b
}
