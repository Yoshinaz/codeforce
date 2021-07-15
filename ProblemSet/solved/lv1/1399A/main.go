package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t, n, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		ans := "YES"
		input := make([]int, 0)
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			input = append(input, a)
		}
		sort.Slice(input, func(i, j int) bool {
			return input[i] < input[j]
		})
		for j := 1; j < n; j++ {
			if abs(input[j]-input[j-1]) > 1 {
				ans = "NO"
			}
		}
		fmt.Println(ans)
	}

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
