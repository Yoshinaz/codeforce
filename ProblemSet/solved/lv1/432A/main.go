package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] += k
		if a[i] <= 5 {
			count++
		}
	}
	fmt.Println(count / 3)
}
