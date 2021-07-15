package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		set := make(map[int]bool, 0)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			set[a] = true
		}
		fmt.Println(len(set))
	}
}
