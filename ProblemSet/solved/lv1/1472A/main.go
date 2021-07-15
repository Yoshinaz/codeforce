package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, w, h int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &w, &h, &n)
		count := 1
		for w%2 == 0 || h%2 == 0 {
			if w%2 == 0 {
				count *= 2
				w /= 2
			}
			if h%2 == 0 {
				count *= 2
				h /= 2
			}
		}
		if count >= n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
