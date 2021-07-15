package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &m, &n)
	data := make([]string, m)
	col := make(map[int]bool, 0)
	row := make(map[int]bool, 0)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &data[i])
		for j := 0; j < n; j++ {
			if data[i][j] == 'S' {
				col[j] = true
				row[i] = true
			}
		}
	}
	fmt.Println(m*n - len(col)*len(row))
}
