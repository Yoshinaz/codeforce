package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	ans := "NO"
	for i := n + 1; i <= m; i++ {
		if IsPrimeSqrt(i) {
			if i == m {
				ans = "YES"
			}
			break
		}
	}
	fmt.Println(ans)
}

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
