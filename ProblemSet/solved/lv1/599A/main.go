package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var d1, d2, d3 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &d1, &d2, &d3)
	d1, d2, d3 = sorta(d1, d2, d3)
	if d1+d2 < d3 {
		fmt.Println((d1 + d2) * 2)
	} else {
		fmt.Println(d1 + d2 + d3)
	}
}

func sorta(a, b, c int) (int, int, int) {
	aa := make([]int, 0)
	aa = append(aa, a)
	aa = append(aa, b)
	aa = append(aa, c)
	sort.Slice(aa, func(i, j int) bool {
		return aa[i] < aa[j]
	})
	return aa[0], aa[1], aa[2]
}
