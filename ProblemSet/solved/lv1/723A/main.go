package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	a := make([]int, 3)
	fmt.Fscan(in, &a[0], &a[1], &a[2])

	m := mid(a)
	dis := 0
	for _, v := range a {
		dis += abs(v - m)
	}
	fmt.Println(dis)
}

func mid(a []int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a[1]
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
