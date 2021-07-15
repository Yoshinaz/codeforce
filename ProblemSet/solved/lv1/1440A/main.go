package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n, c0, c1, h int
	var input string
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &c0, &c1, &h, &input)
		if h+c0 < c1 {
			c1 = h + c0
		}
		if h+c1 < c0 {
			c0 = h + c1 //no problem coz if c1 = h+c0  --> h+c1>c0
		}
		count0 := 0
		for _, v := range input {
			if v == '0' {
				count0++
			}
		}
		count1 := n - count0
		fmt.Println(count0*c0 + count1*c1)
	}
}
