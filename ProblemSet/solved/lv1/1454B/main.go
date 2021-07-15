package main

import (
	"bufio"
	"fmt"
	"os"
)

type bid struct {
	idx   int
	count int
}

func main() {
	var n, t, a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		data := make(map[int]*bid, 0)
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a)
			if v, ok := data[a]; ok {
				v.count++
				//data[a] = bid{v.idx, v.count + 1}
			} else {
				data[a] = &bid{j + 1, 1}
			}
		}
		minBid := n + 1
		ansIdx := -1
		for k, v := range data {
			if minBid > k && v.count == 1 {
				ansIdx = v.idx
				minBid = k
			}
		}
		//fmt.Println(data)
		fmt.Println(ansIdx)

	}
}
