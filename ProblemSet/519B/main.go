package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	a1 := make(map[int]int, n)
	a2 := make(map[int]int, n-1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		a1[a] = a1[a] + 1
	}
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a)
		a2[a] = a2[a] + 1
		a1[a] = a1[a] - 1
	}
	for i := 0; i < n-2; i++ {
		fmt.Fscan(in, &a)
		a2[a] = a2[a] - 1
	}
	ans1 := 0
	for k, v := range a1 {
		if v > 0 {
			ans1 = k
		}
	}
	ans2 := 0
	for k, v := range a2 {
		if v > 0 {
			ans2 = k
		}
	}
	fmt.Println(ans1)
	fmt.Println(ans2)
}
