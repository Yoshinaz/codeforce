package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n, a, la int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &la)
		isStartIncre := false
		count := 1
		for j := 1; j < n; j++ {
			fmt.Fscan(in, &a)
			if a > la {
				isStartIncre = false
			}
			if isStartIncre {
				a++
			} else if a == la {
				isStartIncre = true
				a++
			}
			if a != la {
				count++
			}
			la = a
		}
		fmt.Println(count)
	}
}
