package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for j := 0; j < t; j++ {
		fmt.Fscan(in, &n)
		m := n / 2
		if m%2 == 0 {
			fmt.Println("YES")
			sum := 0
			for i := 1; i <= m; i++ {
				fmt.Printf("%d ", i*2)
				sum += i * 2
			}
			sum2 := 0
			for i := 0; i < m-1; i++ {
				fmt.Printf("%d ", i*2+1)
				sum2 += i*2 + 1
			}
			fmt.Printf("%d\n", sum-sum2)
		} else {
			fmt.Println("NO")
		}
	}
}
