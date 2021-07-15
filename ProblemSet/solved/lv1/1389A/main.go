package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, l, r int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &l, &r)
		if l*2 <= r {
			fmt.Println(l, l*2)
		} else {
			fmt.Println(-1, -1)
		}
	}
}
