package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b, c, d int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		fmt.Println(b, c, c)
	}
}
