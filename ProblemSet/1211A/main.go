package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type prob struct {
	dif int
	idx int
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	p := make([]prob, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].dif)
		p[i].idx = i + 1
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].dif < p[j].dif
	})
	count := 1
	nprob := 3
	ans := make([]prob, nprob)
	ans[0] = p[0]
	for _, v := range p {
		if v.dif != ans[count-1].dif {
			ans[count] = v
			count++
		}
		if count == nprob {
			break
		}
	}
	if count == nprob {
		for i := 0; i < nprob; i++ {
			fmt.Printf("%d ", ans[i].idx)
		}
		fmt.Println()
	} else {
		for i := 0; i < nprob; i++ {
			fmt.Printf("-1 ")
		}
		fmt.Println()
	}

}
