package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var input string
	fmt.Fscan(in, &input)
	isDanger := false
	last := '2'
	count := 1
	for _, v := range input {
		if v == last {
			count++
		} else {
			count = 1
		}
		if count >= 7 {
			isDanger = true
		}
		last = v
	}
	if isDanger {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
