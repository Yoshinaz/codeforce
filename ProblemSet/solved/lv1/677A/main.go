package main

import "fmt"

func main() {
	var n, h, a int
	fmt.Scanf("%d %d\n", &n, &h)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		if a > h {
			count++
		}
	}
	fmt.Println(count + n)
}
