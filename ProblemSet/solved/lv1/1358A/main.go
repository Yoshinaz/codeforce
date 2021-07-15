package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		ans := n / 2 * m
		if n%2 == 1 {
			ans += m / 2
			if m%2 == 1 {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
