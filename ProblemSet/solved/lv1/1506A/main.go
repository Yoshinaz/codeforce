package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var m, n, d int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &m, &n, &d)
		col := d/m + 1
		row := d % m
		if row == 0 {
			row = m
			col -= 1
		}
		ans := (row-1)*n + col
		fmt.Println(ans)
	}
}
