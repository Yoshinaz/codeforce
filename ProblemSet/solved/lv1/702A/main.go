package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	count := 1
	fmt.Fscan(in, &a)
	la := a
	max := 1
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a)
		if a > la {
			count++
			if max < count {
				max = count
			}
		} else {
			count = 1
		}
		la = a

	}
	fmt.Println(max)
}
