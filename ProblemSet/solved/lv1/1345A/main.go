package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, x, y int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &x, &y)
		if x == 1 || y == 1 || (x == 2 && y == 2) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
