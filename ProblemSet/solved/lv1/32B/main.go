package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &s)
	var ans strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			ans.WriteString("0")
		} else if s[i] == '-' {
			i++
			if s[i] == '.' {
				ans.WriteString("1")
			} else {
				ans.WriteString("2")
			}
		}
	}
	fmt.Println(ans.String())
}
