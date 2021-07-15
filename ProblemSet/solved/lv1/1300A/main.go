package main

import (
	"bufio"
	"fmt"
	"os"
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
		sum := 0
		count := 0
		for j := 0; j < n; j++ {
			if data[j] == 0 {
				data[j] = 1
				count++
			}
			sum += data[j]
		}
		if sum == 0 {
			count++
		}
		fmt.Println(count)
	}
}
