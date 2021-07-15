package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		if n <= 3 {
			fmt.Println(n - 1)
		} else {
			fmt.Println(n&1 + 2)
		}
	}
}
