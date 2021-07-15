package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		if a > b {
			count++
		} else if a < b {
			count--
		}
	}
	if count > 0 {
		fmt.Println("Mishka")
	} else if count < 0 {
		fmt.Println("Chris")
	} else {
		fmt.Println("Friendship is magic!^^")
	}
}
