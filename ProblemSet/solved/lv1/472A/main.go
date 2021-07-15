package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	if n%2 == 0 {
		fmt.Println(6, n-6)
	} else {
		fmt.Println(9, n-9)
	}
}
