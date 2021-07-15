package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &s)
		la := -1
		a := -1
		max := 0
		for i, v := range s {
			if v == 'A' {
				a = i
			}
			if la != -1 && max < a-la-1 {
				max = a - la - 1
			}
			la = a
		}
		if la != -1 && max < n-la-1 {
			max = n - la - 1
		}
		fmt.Println(max)
	}
}
