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
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &cards[i])
	}
	left := 0
	right := n - 1
	sum1 := 0
	sum2 := 0
	for i := 0; i < n; i++ {
		chose := 0
		if cards[left] > cards[right] {
			chose = cards[left]
			left++
		} else {
			chose = cards[right]
			right--
		}
		if i%2 == 0 {
			sum1 += chose
		} else {
			sum2 += chose
		}
	}
	fmt.Println(sum1, sum2)
}
