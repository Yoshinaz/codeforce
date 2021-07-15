package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, xn, yn, p int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	q := make([]bool, n+1)
	fmt.Fscan(in, &xn)
	for i := 0; i < xn; i++ {
		fmt.Fscan(in, &p)
		q[p] = true
	}
	fmt.Fscan(in, &yn)
	for i := 0; i < yn; i++ {
		fmt.Fscan(in, &p)
		q[p] = true
	}
	ans := true
	for i := 1; i <= n; i++ {
		ans = ans && q[i]
	}
	if ans {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}
