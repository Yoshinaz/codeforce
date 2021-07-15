package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}
		l := 0
		r := n - 1
		for l < r {
			fmt.Printf("%d ", a[l])
			fmt.Printf("%d ", a[r])
			l++
			r--
		}
		if l == r {
			fmt.Printf("%d", a[l])
		}
		fmt.Println()
	}
}
