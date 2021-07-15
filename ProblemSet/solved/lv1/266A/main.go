package main

import (
	"fmt"
)

func main() {
	var n int
	var stones string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &stones)
	last := stones[0]
	ns := len(stones)
	count := 0
	for i := 1; i < ns; i++ {
		if last == stones[i] {
			count++
		}
		last = stones[i]
	}
	fmt.Println(count)
}
