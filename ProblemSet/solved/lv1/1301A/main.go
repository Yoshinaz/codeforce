package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var a, b, c string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c)
		ans := "YES"
		for j := 0; j < len(a); j++ {
			if !checkAns(a[j], b[j], c[j]) {
				ans = "NO"
				break
			}
		}
		fmt.Println(ans)
	}
}

func checkAns(a, b, c byte) bool {
	if a == c {
		return true
	}
	if b == c {
		return true
	}
	return false
}
