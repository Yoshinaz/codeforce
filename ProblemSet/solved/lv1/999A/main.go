package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &k)
	canDo := true
	count := 0
	count2 := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		count2++
		if a > k {
			canDo = false
			count2 = 0
		}
		if canDo {
			count++
		}
	}
	if canDo {
		count2 = 0
	}
	fmt.Println(count + count2)
}
