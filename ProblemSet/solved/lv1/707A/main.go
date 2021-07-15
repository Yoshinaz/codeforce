package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	var a string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	isColor := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a)
			if a != "W" && a != "B" && a != "G" {
				isColor = true
			}
		}
	}
	if isColor {
		fmt.Println("#Color")
	} else {
		fmt.Println("#Black&White")
	}
}
