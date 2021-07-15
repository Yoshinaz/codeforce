package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s1, s2 string
	var ans strings.Builder
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &s1)
	fmt.Fscan(in, &s2)

	for i := 0; i < len(s1); i++ {
		if s2[i] == s1[i] {
			ans.WriteString("0")
		} else {
			ans.WriteString("1")
		}
	}
	fmt.Println(ans.String())
}
