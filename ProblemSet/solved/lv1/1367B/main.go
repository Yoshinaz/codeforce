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
		c1 := 0
		c2 := 0
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			if j%2 != a%2 {
				if j%2 == 1 {
					c1++
				} else {
					c2++
				}
			}
		}
		if c1 == c2 {
			fmt.Println(c1)
		} else {
			fmt.Println(-1)
		}
	}

}
