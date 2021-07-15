package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		mul := 1
		fmt.Fscan(in, &n)
		ans := make([]int, 0)
		for n > 0 {
			x := n % 10
			if x != 0 {
				ans = append(ans, x*mul)
			}
			n /= 10
			mul *= 10
		}
		fmt.Println(len(ans))
		for _, v := range ans {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
