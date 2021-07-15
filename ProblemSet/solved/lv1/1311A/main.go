package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		if a < b {
			if b%2 == a%2 {
				fmt.Println(2)
			} else {
				fmt.Println(1)
			}
		} else if a > b {
			if b%2 == a%2 {
				fmt.Println(1)
			} else {
				fmt.Println(2)
			}
		} else {
			fmt.Println(0)
		}
	}
}
