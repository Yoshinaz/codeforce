package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		count := make(map[string]int, 0)
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &s)
			for _, v := range s {
				count[string(v)]++
			}
		}
		ans := "YES"
		for _, v := range count {
			if v%n != 0 {
				ans = "NO"
				break
			}
		}
		fmt.Println(ans)
	}

}
