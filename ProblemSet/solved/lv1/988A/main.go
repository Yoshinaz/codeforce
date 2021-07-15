package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, n, a int
	in := bufio.NewReader(os.Stdin)
	m := make(map[int]int, 0)
	fmt.Fscan(in, &k, &n)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a)
		m[a] = i + 1
	}
	if len(m) >= n {
		fmt.Println("YES")
		for _, v := range m {
			fmt.Printf("%d ", v)
			n--
			if n == 0 {
				break
			}
		}
		fmt.Println()
	} else {
		fmt.Println("NO")
	}
}
