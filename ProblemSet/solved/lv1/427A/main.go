package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, c int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	p := 0
	count := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c)
		p = p + c
		if p < 0 {
			count++
			p = 0
		}
	}
	fmt.Println(count)
}
