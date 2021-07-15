package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var k, d int64
	var sign string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &sign, &d)
		if sign == "+" {
			k += d
		} else {
			if d > k {
				count++
			} else {
				k -= d
			}
		}
	}
	fmt.Println(k, count)
}
