package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	a := make([]int, 4)
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a[0], &a[1], &a[2], &a[3])
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sum1 := 0
	sum2 := 0
	for i := 3; i >= 0; i-- {
		if sum1 < sum2 {
			sum1 += a[i]
		} else {
			sum2 += a[i]
		}
	}
	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
