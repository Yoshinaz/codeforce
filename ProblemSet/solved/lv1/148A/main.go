package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, l, m, n, d int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &k)
	fmt.Fscan(in, &l)
	fmt.Fscan(in, &m)
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &d)
	check := make([]bool, d+1)
	for i := k; i <= d; i += k {
		check[i] = true
	}
	for i := l; i <= d; i += l {
		check[i] = true
	}
	for i := m; i <= d; i += m {
		check[i] = true
	}
	for i := n; i <= d; i += n {
		check[i] = true
	}
	ans := 0
	for i := 1; i <= d; i++ {
		if check[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
