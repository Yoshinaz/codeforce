package main

import "fmt"

func main() {
	var n, p, q int
	fmt.Scanf("%d\n", &n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &p, &q)
		if q-p > 1 {
			count++
		}
	}
	fmt.Println(count)
}
