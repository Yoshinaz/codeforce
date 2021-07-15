package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	sum = sum * 2 / n
	for i := 0; i < n; i++ {
		if a[i] != -1 {
			for j := i + 1; j < n; j++ {
				if a[j] == sum-a[i] {
					fmt.Println(i+1, j+1)
					a[j] = -1
					break
				}
			}
		}
	}
}
