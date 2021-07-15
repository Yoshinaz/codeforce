package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	var ans strings.Builder
	ans.WriteString("I hate ")
	for i := 1; i < n; i++ {
		ans.WriteString("that ")
		if i%2 == 0 {
			ans.WriteString("I hate ")
		} else {
			ans.WriteString("I love ")
		}
	}
	ans.WriteString("it")
	fmt.Println(ans.String())
}
