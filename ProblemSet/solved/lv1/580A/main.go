package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, a int
	fmt.Fscan(in, &n)
	count := 0
	la := -1
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		if a >= la {
			count++
		} else {
			count = 1
		}
		if count > ans {
			ans = count
		}
		la = a
	}
	fmt.Println(ans)
}
