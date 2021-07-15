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
		fmt.Fscan(in, &n, &a)
		la := a
		hasDif := false
		for j := 1; j < n; j++ {
			fmt.Fscan(in, &a)
			if a != la {
				hasDif = true
			}
		}
		if hasDif {
			fmt.Println(1)
		} else {
			fmt.Println(n)
		}
	}
}
