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
		aa := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &aa[j])
		}
		for j := n - 1; j >= 0; j-- {
			fmt.Printf("%d ", aa[j])
		}
		fmt.Println()
	}
}
