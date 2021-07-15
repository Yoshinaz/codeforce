package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &s)
	ss := strings.ToLower(s)
	count := make(map[string]int)
	for i := 0; i < n; i++ {
		count[string(ss[i])]++
	}
	c := 0
	for _, _ = range count {
		c++
	}
	if c == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
