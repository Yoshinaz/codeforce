package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Println(int64(math.Ceil(float64(n+1) / 2.0)))

	}
}
