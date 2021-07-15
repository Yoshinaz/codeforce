package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, h int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	f := make([]int, n+1)
	sum := 0
	f[0] = 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h)
		sum += h
		f[i+1] = sum
	}
	min := f[k] - f[0]
	ans := 1
	for i := 1; i <= n-k; i++ {
		tmp := f[i+k] - f[i]
		if tmp < min {
			min = tmp
			ans = i + 1
		}
	}
	fmt.Println(ans)
}
