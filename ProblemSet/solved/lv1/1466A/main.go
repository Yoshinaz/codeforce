package main

import (
	"bufio"
	"fmt"
	"os"
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
		set := make(map[int]bool, 0)
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				set[a[k]-a[j]] = true
			}
		}
		fmt.Println(len(set))
	}
}
