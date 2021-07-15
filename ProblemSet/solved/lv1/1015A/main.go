package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, l, r int
	var flag [100]bool
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l, &r)
		for j := l - 1; j < r; j++ {
			flag[j] = true
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		if !flag[i] {
			count++
		}
	}
	fmt.Println(count)
	for i := 0; i < m; i++ {
		if !flag[i] {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}
