package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	var s string
	tmp := "2020"
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &s)
		l := 0
		r := 3
		if n <= r {
			fmt.Println("NO")
		} else {
			for j := 0; j <= r; j++ {
				if tmp[l] != s[j] {
					break
				}
				l++
			}
			for j := n - 1; j > 0 && r >= 0; j-- {
				if tmp[r] != s[j] {
					break
				}
				r--
			}
			if l > r {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
