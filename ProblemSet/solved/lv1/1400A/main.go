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
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &s)
		idx := len(s) / 2
		for j := 0; j < n; j++ {
			fmt.Printf("%s", string(s[idx]))
		}
		fmt.Println()
	}
}
