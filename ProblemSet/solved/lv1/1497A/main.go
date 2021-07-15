package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}
		sort.Ints(a)
		tmp := make([]int, 0)
		fmt.Printf("%d ", a[0])
		for j := 1; j < n; j++ {
			if a[j] != a[j-1] {
				fmt.Printf("%d ", a[j])
			} else {
				tmp = append(tmp, a[j])
			}
		}
		for _, v := range tmp {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
