package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		x := a | b
		fmt.Println((a ^ x) + (b ^ x))
	}
}
