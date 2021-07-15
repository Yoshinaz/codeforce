package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, x, y, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &x, &y, &a, &b)
		d := y - x
		if d%(a+b) == 0 {
			fmt.Println(d / (a + b))
		} else {
			fmt.Println(-1)
		}
	}
}
