package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var a [3]int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a[0], &a[1], &a[2])
		m := max(a[0], a[1])
		m = max(m, a[2])
		min := a[0]
		count := 0
		for j := 0; j < 3; j++ {
			if m == a[j] {
				count++
			} else {
				min = a[j]
			}
		}
		if count >= 2 {
			fmt.Println("YES")
			fmt.Printf("%d %d %d\n", m, min, min)
		} else {
			fmt.Println("NO")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
