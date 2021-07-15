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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	max := a[0]
	for i := 1; i < n; i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		count += max - a[i]
	}
	fmt.Println(count)
}
