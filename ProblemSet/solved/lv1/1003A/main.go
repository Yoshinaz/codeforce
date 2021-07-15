package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	coins := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		coins[a]++
	}
	max := 0
	for _, v := range coins {
		if max < v {
			max = v
		}
	}
	fmt.Println(max)
}
