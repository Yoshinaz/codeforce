package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, a, b, c int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &a)
		fmt.Fscan(in, &b)
		for j := 0; j < n-2; j++ {
			fmt.Fscan(in, &c)
		}

		if a+b <= c {
			fmt.Println(1, 2, n)
		} else {
			fmt.Println(-1)
		}
	}
}
