package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, c, s1, s2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &c)
	fmt.Fscan(in, &s1)
	count := 1
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &s2)
		if s2-s1 <= c {
			count++
		} else {
			count = 1
		}
		s1 = s2
	}
	fmt.Println(count)
}
