package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &s)
	ans := ""
	j := 0
	for i := 0; i < n; i += j {
		j += 1
		ans += string(s[i])
	}
	fmt.Println(ans)
}
