package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type laptop struct {
	p int
	q int
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	lap := make([]laptop, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &lap[i].p, &lap[i].q)
	}
	if isExist(lap, n) {
		fmt.Println("Happy Alex")
	} else {
		fmt.Println("Poor Alex")
	}
}

func isExist(lap []laptop, n int) bool {
	sort.Slice(lap, func(i, j int) bool {
		return lap[i].p < lap[j].p
	})

	// laptop can be same price
	maxqtmp := lap[0].q
	maxp := lap[0].p
	maxq := lap[0].q

	for i := 1; i < n; i++ {
		if lap[i].q < maxq && lap[i].p > maxp {
			return true
		} else if lap[i].q > maxq && lap[i].p == maxp {
			maxqtmp = lap[i].q
		} else {
			maxq = maxqtmp
			maxp = lap[i].p
			if lap[i].q > maxq {
				maxq = lap[i].q
			}
		}
	}
	return false
}
