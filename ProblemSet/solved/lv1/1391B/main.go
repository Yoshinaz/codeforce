package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, m int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		ans := 0
		for j := 0; j < n-1; j++ {
			fmt.Fscan(in, &s)
			if s[m-1] == 'R' {
				ans++
			}
		}
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			if s[j] == 'D' {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
