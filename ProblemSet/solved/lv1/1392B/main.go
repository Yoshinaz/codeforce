package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	var k int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &k)
		data := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &data[j])
		}
		if k%2 == 0 {
			do(data)
			do(data)
		} else {
			do(data)
		}
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", data[j])
		}
		fmt.Println()
	}
}

func do(a []int) {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		a[i] = max - a[i]
	}
}
