package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		ans := make([]int, 0)
		tmp := make(map[int]bool, 0)
		for j := 0; j < n*2; j++ {
			fmt.Fscan(in, &a)
			if !tmp[a] {
				tmp[a] = true
				ans = append(ans, a)
			}
		}
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", ans[j])
		}
		fmt.Println()
	}
}
