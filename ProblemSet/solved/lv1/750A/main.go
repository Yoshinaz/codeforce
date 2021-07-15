package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &t)
	count := 0
	time := 240 - t
	for i := 1; i <= n && time >= 0; i++ {
		time = time - 5*i
		if time >= 0 {
			count++
		}
	}
	fmt.Println(count)
}
