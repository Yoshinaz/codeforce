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
		fmt.Println("Mahmoud")
	} else {
		fmt.Println("Ehab")
	}
}
