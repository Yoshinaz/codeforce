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
		reminder := x % y
		if reminder == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(y - x%y)
		}
	}
}
