package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, m, a int
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		data := make(map[int]bool, 0)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a)
			data[a] = true
		}
		hasAns := false
		ans := -1
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &a)
			if data[a] {
				hasAns = true
				ans = a
			}
		}
		if hasAns {
			fmt.Println("YES")
			fmt.Println(1, ans)
		} else {
			fmt.Println("NO")
		}
	}
}
