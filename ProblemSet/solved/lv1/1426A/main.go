package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t, x, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &x)
		fmt.Printf("%.0f\n", math.Ceil(float64(pos(n-2))/float64(x))+1)
	}
}

func pos(a int) int {
	if a < 0 {
		return 0
	}
	return a
}
