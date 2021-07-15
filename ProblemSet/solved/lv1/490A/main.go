package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	var student [3][]int
	for i := 0; i < 3; i++ {
		student[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		student[a-1] = append(student[a-1], i+1)
	}
	m := min(len(student[0]), len(student[1]))
	m = min(m, len(student[2]))
	fmt.Println(m)
	for i := 0; i < m; i++ {
		fmt.Println(student[0][i], student[1][i], student[2][i])
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
