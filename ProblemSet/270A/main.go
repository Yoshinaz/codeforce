package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, x int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &x)
		if 360%(180-x) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
