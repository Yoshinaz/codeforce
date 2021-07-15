package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type data struct {
	price int
	stock int
}

func main() {
	var t, bun, beef, chick, bp, cp int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &bun, &beef, &chick, &bp, &cp)
		bun = bun / 2
		goods := make([]data, 0)
		goods = append(goods, data{bp, beef})
		goods = append(goods, data{cp, chick})
		sort.Slice(goods, func(i, j int) bool {
			return goods[i].price > goods[j].price
		})
		profit := 0
		for i := 0; i < len(goods); i++ {
			if bun > goods[i].stock {
				bun -= goods[i].stock
				profit += goods[i].stock * goods[i].price
			} else {
				profit += goods[i].price * bun
				bun = 0
			}
		}
		fmt.Println(profit)
	}
}

func swap(a, b, c, d int) (int, int, int, int) {
	if a > b {
		return a, b, c, d
	} else {
		return b, a, d, c
	}
}
