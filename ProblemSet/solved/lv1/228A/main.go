package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s int
	shoe := make(map[int]bool)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		fmt.Fscan(in, &s)
		shoe[s] = true
	}
	ans := 0
	for _, _ = range shoe {
		ans++
	}
	fmt.Println(4 - ans)
}
