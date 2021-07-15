package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t int
	var n, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		b := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &b[j])
		}
		sort.Slice(a, func(i, j int) bool {
			return a[i] > a[j]
		})
		sort.Slice(b, func(i, j int) bool {
			return b[i] > b[j]
		})
		sum := 0
		ai := 0
		bi := 0
		for j := 0; j < n; j++ {
			if a[ai] >= b[bi] || k == 0 {
				sum += a[ai]
				ai++
			} else {
				sum += b[bi]
				bi++
				k--
			}
		}
		fmt.Println(sum)
	}
}
