package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &m, &n)
	fmt.Println(count(m, n))

}

func count(m, n int) int {
	count := 0
	for a := 0; a <= m; a++ {
		for b := 0; b <= n; b++ {
			if a*a+b == n && a+b*b == m {
				count++
			}
		}
	}
	return count
}
