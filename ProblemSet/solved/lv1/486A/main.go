package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	even := n / 2
	odd := n - even
	ans := (even)*(even+1) - odd*odd
	fmt.Println(ans)
}
