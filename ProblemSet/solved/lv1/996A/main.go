package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	bil := []int{100, 20, 10, 5}
	in := bufio.NewReader(os.Stdin)
	ans := 0
	fmt.Fscan(in, &n)
	for _, v := range bil {
		ans = ans + n/v
		n = n % v
	}
	ans = ans + n
	fmt.Println(ans)
}
