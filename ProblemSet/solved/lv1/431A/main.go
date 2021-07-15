package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a [4]int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a[0], &a[1], &a[2], &a[3])
	fmt.Fscan(in, &s)
	count := 0
	for _, v := range s {
		idx, _ := strconv.Atoi(string(v))
		count += a[idx-1]
	}
	fmt.Println(count)
}
