package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var a, b, c int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c)
		fmt.Println(a + b + c - 1)
	}
}
