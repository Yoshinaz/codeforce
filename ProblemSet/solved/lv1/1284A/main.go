package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, q, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	s := make([]string, n)
	t := make([]string, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &t[i])
	}
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &a)
		fmt.Println(s[(a-1)%n] + t[(a-1)%m])
	}
}
