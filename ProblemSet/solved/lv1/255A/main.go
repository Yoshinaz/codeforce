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
	chest := 0
	bicep := 0
	back := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		if i%3 == 0 {
			chest += a
		} else if i%3 == 1 {
			bicep += a
		} else {
			back += a
		}
	}
	max, ans := findans(chest, bicep, "chest", "biceps")
	max, ans = findans(back, max, "back", ans)
	fmt.Println(ans)
}

func findans(a, b int, s1, s2 string) (int, string) {
	if a > b {
		return a, s1
	}
	return b, s2
}
