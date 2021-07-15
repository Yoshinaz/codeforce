package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)
	fmt.Println(len(s) + 1)
}
