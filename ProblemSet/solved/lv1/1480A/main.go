package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	var s string
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &s)
		for i, v := range s {
			if i%2 == 0 {
				if v == 'a' {
					fmt.Printf("b")
				} else {
					fmt.Printf("a")
				}
			} else {
				if v == 'z' {
					fmt.Printf("y")
				} else {
					fmt.Printf("z")
				}
			}
		}
		fmt.Println()
	}
}
