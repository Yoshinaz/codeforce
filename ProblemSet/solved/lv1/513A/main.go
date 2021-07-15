package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n1, n2, k1, k2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n1, &n2, &k1, &k2)
	if n1 <= n2 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
