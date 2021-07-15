package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t, k int
	var a, b, c, d float64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c, &d, &k)
		aa := math.Ceil(a / c)
		bb := math.Ceil(b / d)
		if int(aa)+int(bb) <= k {
			fmt.Println(aa, bb)
		} else {
			fmt.Println(-1)
		}
	}
}
