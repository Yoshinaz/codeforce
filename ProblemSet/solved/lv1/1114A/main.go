package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x, y, z, a, b, c int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &x, &y, &z, &a, &b, &c)
	ans := findAns(x, y, z, a, b, c)
	fmt.Println(ans)
}

func findAns(x, y, z, a, b, c int) string {
	if x > a {
		return "NO"
	}
	a = a - x
	if y > a+b {
		return "NO"
	}
	if y+z > a+b+c {
		return "NO"
	}
	return "YES"
}
