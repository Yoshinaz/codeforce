package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		if n > 2 {
			fmt.Println(m * 2)
		} else if n == 1 {
			fmt.Println(0)
		} else {
			fmt.Println(m)
		}
	}
}
