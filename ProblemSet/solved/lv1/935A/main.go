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
	fact := make([]int, 0)
	count := 0
	for n%2 == 0 {
		count++
		n /= 2
	}
	fact = append(fact, count)
	count = 0
	for i := 3; n > 1; i += 2 {
		for n%i == 0 {
			n /= i
			count++
		}
		if count > 0 {
			fact = append(fact, count)
			count = 0
		}
	}
	ans := 1
	for i := 0; i < len(fact); i++ {
		ans *= (fact[i] + 1)
	}
	fmt.Println(ans - 1)
}
