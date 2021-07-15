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
	la := 0
	count := 0
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		if a <= la {
			count++
			ans = append(ans, la)
		}
		la = a
	}
	count++
	ans = append(ans, a)
	fmt.Println(count)
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
