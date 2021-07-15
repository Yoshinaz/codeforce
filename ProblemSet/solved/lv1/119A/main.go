package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a, &b, &n)
	count := 0
	for n > 0 {
		if count%2 == 0 {
			m = gcd(n, a)
		} else {
			m = gcd(n, b)
		}
		count++
		n -= m
	}
	fmt.Println((count + 1) % 2)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}