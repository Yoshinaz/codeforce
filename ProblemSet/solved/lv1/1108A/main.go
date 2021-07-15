package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, l1, r1, l2, r2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &l1, &r1, &l2, &r2)
		if l1 != r2 {
			fmt.Println(l1, r2)
		} else {
			fmt.Println(r1, l2)
		}
	}
}
