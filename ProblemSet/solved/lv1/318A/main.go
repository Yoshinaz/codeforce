package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var k, n int64
	fmt.Fscan(in, &k, &n)
	if n <= (k+1)/2 {
		fmt.Println(2*n - 1)
	} else {
		n -= (k + 1) / 2
		fmt.Println(2 * n)
	}
}
