package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t, n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		sum := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			sum += a
		}
		fmt.Printf("%.0f\n", math.Ceil(float64(sum)/float64(n)))
	}
}
