package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}
		min := findMin(a)
		c := countMin(a, min)
		fmt.Println(n - c)
	}
}

func findMin(a []int) int {
	min := math.MaxInt16
	n := len(a)
	for j := 0; j < n; j++ {
		if min > a[j] {
			min = a[j]
		}
	}
	return min
}

func countMin(a []int, min int) int {
	c := 0
	n := len(a)
	for j := 0; j < n; j++ {
		if min == a[j] {
			c++
		}
	}
	return c
}
