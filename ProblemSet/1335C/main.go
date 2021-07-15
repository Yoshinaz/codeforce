package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		m := make(map[int]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			m[a]++
		}
		maxSame := 0
		maxDif := 0
		for _, v := range m {
			if v > maxSame {
				maxSame = v
			}
			maxDif++
		}
		if maxSame == maxDif {
			fmt.Println(maxSame - 1)
		} else {
			fmt.Println(min(maxSame, maxDif))
		}
	}

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
