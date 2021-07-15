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
		fmt.Fscan(in, &n)
		if n <= 30 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			if n == 36 {
				fmt.Println(6, 10, 15, 5)
			} else if n == 40 {
				fmt.Println(6, 10, 15, 9)
			} else if n == 44 {
				fmt.Println(6, 10, 15, 13)
			} else {
				fmt.Println(6, 10, 14, n-30)
			}
		}
	}
}
