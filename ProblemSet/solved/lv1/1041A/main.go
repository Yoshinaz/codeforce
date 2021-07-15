package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	var a int64
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		if max < a {
			max = a
		}
		if min > a {
			min = a
		}
	}
	fmt.Println(max - min + int64(1) - int64(n))
}
