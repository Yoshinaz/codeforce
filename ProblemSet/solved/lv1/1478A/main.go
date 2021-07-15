package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n, a, la int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &la)
		count := 0
		ans := 0
		for j := 1; j < n; j++ {
			fmt.Fscan(in, &a)
			if a == la {
				count++
			} else {
				if count > ans {
					ans = count
				}
				la = a
				count = 0
			}
		}
		if count > ans {
			ans = count
		}

		fmt.Println(ans + 1)
	}
}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	var t, n, a int
// 	fmt.Fscan(in, &t)
// 	for i := 0; i < t; i++ {
// 		fmt.Fscan(in, &n)
// 		bucket := make([]int, n)
// 		for j := 0; j < n; j++ {
// 			fmt.Fscan(in, &a)
// 			bucket[a-1]++
// 		}
// 		ans := 0
// 		for j := 0; j < n; j++ {
// 			if ans < bucket[j] {
// 				ans = bucket[j]
// 			}
// 		}

// 		fmt.Println(ans)
// 	}
// }
