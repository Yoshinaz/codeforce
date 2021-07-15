package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scanf("%d\n", &n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d\n", &a, &b, &c)
		if a+b+c > 1 {
			count++
		}
	}
	fmt.Println(count)
}
