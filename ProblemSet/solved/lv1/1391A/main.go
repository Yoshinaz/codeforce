package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", j+1)
		}
		fmt.Println()
	}
}
