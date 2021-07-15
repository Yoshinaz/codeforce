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
	i := 1
	sum := 0
	sum2 := 0
	level := 0
	for i = 1; sum2 < n; i++ {
		sum += i
		sum2 += sum
		if sum2 > n {
			break
		}
		level++
	}
	fmt.Println(level)
}
