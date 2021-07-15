package main

import (
	"bufio"
	"fmt"
	"os"
)

type distinct struct {
	gang int
	idx  int
}

func main() {
	var t, n, x int
	//file, _ := os.OpenFile("tmp", os.O_RDONLY, os.ModePerm)
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			pira = append(pira, x)
		}
		maxIL := -2
		maxL := 0
		maxIR := -2
		maxR := 0
		for i := 1; i < n; i++ {
			if maxL < pira[i] && pira[i-1] < pira[i] {
				maxIL = i
				maxL = pira[i]
			}
		}
		for i := n - 1; i > 0; i-- {
			if maxR < pira[i-1] && pira[i] < pira[i-1] {
				maxIR = i - 1
				maxR = pira[i-1]
			}
		}
		if maxR > maxL {
			fmt.Println(maxIR + 1)
		} else {
			fmt.Println(maxIL + 1)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
