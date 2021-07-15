package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, h, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	time := 24 * 60
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &h, &m)
		fmt.Println(time - h*60 - m)
	}
}
