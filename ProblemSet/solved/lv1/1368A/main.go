package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &n)
		a, b := minmax(a, b)
		count := 0
		for b <= n {
			//fmt.Println(a, b)
			t := b
			b = a + b
			a = t
			count++

		}
		fmt.Println(count)
	}

}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
