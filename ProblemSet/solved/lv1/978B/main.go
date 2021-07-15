package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &s)
	count := 0
	ans := 0
	bx := rune('x')
	for _, v := range s {
		if v == bx {
			count++
		} else {
			count = 0
		}
		if count >= 3 {
			ans++
		}
	}
	fmt.Println(ans)
}
