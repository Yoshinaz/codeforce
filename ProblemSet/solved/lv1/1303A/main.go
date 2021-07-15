package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &s)
		s = removeLeadingZero(s)
		count := 0
		ans := 0
		for _, v := range s {
			if v == '1' {
				ans += count
				count = 0
			} else {
				count++
			}
		}
		fmt.Println(ans)
	}
}

func removeLeadingZero(s string) string {
	idx := 0
	for ; idx < len(s) && s[idx] == '0'; idx++ {
	}
	return s[idx:]
}
