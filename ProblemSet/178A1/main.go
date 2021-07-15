package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		p := getpowerof2(n - i - 1)
		ans += a[i]
		a[i+p] += a[i]

		fmt.Println(ans)
	}
}

func getpowerof2(n int) int {
	po := 1
	for po*2 <= n {
		po *= 2
	}
	return po
}
