package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		data := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &data[j])
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
		ans := 1
		for j := 0; j < n-1; j++ {
			if data[j+1]-data[j] == 1 {
				ans = 2
			}
		}
		fmt.Println(ans)
	}
}
