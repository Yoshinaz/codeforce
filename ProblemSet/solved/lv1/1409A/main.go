package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		fmt.Printf("%.0f\n", math.Ceil(float64(abs(a-b))/10.0))
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
