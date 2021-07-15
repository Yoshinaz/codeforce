package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	var coin [30]int64
	w := int64(2)
	for i := 0; i < 30; i++ {
		coin[i] = w
		w = w * 2
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m)
		if m == 2 {
			fmt.Println(2)
			continue
		}
		sum1 := int64(0)
		sum2 := int64(0)
		for j := 0; j < m/2-1; j++ {
			sum1 += coin[j]
		}
		for j := m/2 - 1; j < m-1; j++ {
			sum2 += coin[j]
		}
		sum1 += coin[m-1]
		fmt.Println(sum1 - sum2)
	}
}
