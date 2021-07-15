package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type shop struct {
	price int
	n     int
}

func main() {
	var n, m int
	//file, _ := os.OpenFile("in.txt", os.O_RDONLY, os.ModePerm)
	//in := bufio.NewReader(file)
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	price := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &price[i])
	}
	sort.Ints(price)
	shop := toShop(price)
	//fmt.Println(shop)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m)
		idx := bs(shop, m)
		fmt.Println(shop[idx-1].n)
	}
}
func toShop(price []int) []shop {
	s := make([]shop, 0)
	bp := -1
	for i, v := range price {
		if v != bp {
			s = append(s, shop{bp, i})
			bp = v
		}
	}
	s = append(s, shop{bp, len(price)})
	return s
}
func bs(a []shop, key int) int {
	l := 0
	r := len(a) - 1
	m := (l + r) / 2
	for l <= r {
		m = (l + r) / 2
		if key < a[m].price {
			r = m - 1
		} else if key > a[m].price {
			l = m + 1
		} else {
			return m + 1
		}
	}
	return r + 1
}
