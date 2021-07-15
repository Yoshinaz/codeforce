package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t, n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		aa := make([]int, 0)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			aa = append(aa, a)
		}
		sort.Slice(aa, func(i, j int) bool {
			return aa[i] < aa[j]
		})
		min := aa[n-1] - aa[0]
		for j := 0; j < n-1; j++ {
			if min > aa[j+1]-aa[j] {
				min = aa[j+1] - aa[j]
			}
		}
		fmt.Println(min)
	}
}
