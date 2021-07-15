package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a1, a2, a3, b1, b2, b3, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a1, &a2, &a3, &b1, &b2, &b3, &n)
	ans := math.Ceil(float64(a1+a2+a3)/5.0) + math.Ceil(float64(b1+b2+b3)/10.0)
	if int(ans) <= n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
