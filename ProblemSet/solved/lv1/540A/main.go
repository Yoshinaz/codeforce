package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var s1, s2 string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s1, &s2)
	sum := 0
	for i := 0; i < n; i++ {
		i1, _ := strconv.Atoi(string(s1[i]))
		i2, _ := strconv.Atoi(string(s2[i]))
		x := abs(i1 - i2)
		sum += min(x, 10-x)
	}
	fmt.Println(sum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
