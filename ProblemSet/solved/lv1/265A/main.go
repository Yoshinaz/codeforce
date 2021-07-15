package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s1, s2 string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &s1, &s2)
	i := 0

	for j := 0; j < len(s2); j++ {
		if s1[i] == s2[j] {
			i++
		}
	}

	fmt.Println(i + 1)
}
