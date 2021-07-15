package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	set := make(map[int]bool)
	ans := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		if _, ok := set[a[i]]; !ok {
			set[a[i]] = true
			ans = append(ans, a[i])
		}
	}
	fmt.Println(len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Println()
}
