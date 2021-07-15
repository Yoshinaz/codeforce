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
	var x [10][10]int
	for i := 0; i < n; i++ {
		x[i][0], x[0][i] = 1, 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			x[i][j] = x[i-1][j] + x[i][j-1]
		}
	}
	fmt.Println(x[n-1][n-1])
}
