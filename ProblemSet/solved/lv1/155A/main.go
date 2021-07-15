package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, s int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)
	max, min := s, s
	count := 0
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &s)
		if s < min {
			count++
			min = s
		}
		if s > max {
			count++
			max = s
		}
	}
	fmt.Println(count)
}
