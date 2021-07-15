package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	if n%2 == 1 {
		fmt.Println(-1)
	} else {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				fmt.Printf("%d ", i+2)
			} else {
				fmt.Printf("%d ", i)
			}
		}
	}
}
