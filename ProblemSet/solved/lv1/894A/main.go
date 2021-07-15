package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &s)
	countQ := 0
	for _, v := range s {
		if v == 'Q' {
			countQ++
		}
	}
	fq := 0
	ans := 0
	for _, v := range s {
		if v == 'Q' {
			fq++
		}
		if v == 'A' {
			ans += fq * (countQ - fq)
		}
	}
	fmt.Println(ans)
}
