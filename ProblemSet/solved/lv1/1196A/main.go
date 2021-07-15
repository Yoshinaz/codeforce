package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t int
	var a, b, c int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		aa := make([]int64, 0)
		fmt.Fscan(in, &a, &b, &c)
		aa = append(aa, a)
		aa = append(aa, b)
		aa = append(aa, c)
		sort.Slice(aa, func(i, j int) bool {
			return aa[i] < aa[j]
		})
		left := aa[2] - aa[1] + aa[0]
		fmt.Println(left/2 + aa[1])
	}
}
