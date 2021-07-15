package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, t, w int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		sum := 0
		count2 := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &w)
			sum += w
			if w%2 == 0 {
				count2++
			}
		}
		if sum%2 == 0 {
			if count2%2 == 0 || count2 != n {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			fmt.Println("NO")
		}
	}
}
