package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var input string
	fmt.Fscan(in, &input)
	tmp := "HQ9"
	isExecutable := false
	for _, v := range tmp {
		i := strings.Index(input, string(v))
		if i > -1 {
			isExecutable = true
			break
		}
	}
	if isExecutable {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
