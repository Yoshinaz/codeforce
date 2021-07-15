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
	countLetter := make(map[rune]int)
	for _, v := range s {
		countLetter[v]++
	}
	fmt.Fscan(in, &s)
	for _, v := range s {
		countLetter[v]++
	}
	fmt.Fscan(in, &s)
	for _, v := range s {
		countLetter[v]--
	}
	count := 0
	for _, v := range countLetter {
		if v != 0 {
			count++
		}
	}
	if count > 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
