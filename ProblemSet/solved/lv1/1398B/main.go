package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &s)
		count := make([]int, 0)
		c := 0
		for _, v := range s {
			if v == '1' {
				c++
			} else if c != 0 {
				count = append(count, c)
				c = 0
			}
		}
		if c != 0 {
			count = append(count, c)
			c = 0
		}
		sort.Slice(count, func(i, j int) bool {
			return count[i] > count[j]
		})
		ans := 0
		for j := 0; j < len(count); j += 2 {
			ans += count[j]
		}
		fmt.Println(ans)
	}
}
