package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m, &k)
	if n <= m && n <= k {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
