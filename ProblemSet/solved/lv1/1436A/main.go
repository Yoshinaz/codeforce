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
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			m -= a
		}
		if m == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
