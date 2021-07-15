package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	var input string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &input)
		count := 0
		for j := n - 1; j >= 0; j-- {
			if input[j] == ')' {
				count++
			} else {
				break
			}
		}
		if count > n/2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
