package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a)
	r := a % 4
	if r == 0 || r == 3 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
