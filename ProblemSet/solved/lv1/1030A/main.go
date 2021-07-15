package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, p int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	ans := "EASY"
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p)
		if p == 1 {
			ans = "HARD"
		}
	}
	fmt.Println(ans)
}
