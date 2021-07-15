package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	max1 := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		if max1 < a {
			max1 = a
		}
	}
	fmt.Fscan(in, &n)
	max2 := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		if max2 < a {
			max2 = a
		}
	}
	fmt.Println(max1, max2)
}
