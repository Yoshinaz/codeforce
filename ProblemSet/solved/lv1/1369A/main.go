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
		if n%4 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
