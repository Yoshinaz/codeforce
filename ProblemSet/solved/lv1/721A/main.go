package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &s)
	count := 0
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == 'B' {
			count++
		} else {
			if count != 0 {
				ans = append(ans, count)
			}
			count = 0
		}
	}
	if count != 0 {
		ans = append(ans, count)
	}
	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
