package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, h, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	host := make(map[int]int, n)
	guest := make([]int, a)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h, &a)
		host[h] = host[h] + 1
		guest = append(guest, a)
	}
	ans := 0
	for _, v := range guest {
		ans += host[v]
	}
	fmt.Println(ans)
}
