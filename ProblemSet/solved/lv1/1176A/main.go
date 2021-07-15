package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var n int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		count := 0
		for {
			if n == 1 {
				break
			} else if n%2 == 0 {
				n = n / 2
			} else if n%3 == 0 {
				n = 2 * n / 3
			} else if n%5 == 0 {
				n = 4 * n / 5
			} else {
				break
			}
			count++
		}
		if n == 1 {
			fmt.Println(count)
		} else {
			fmt.Println(-1)
		}
	}
}
