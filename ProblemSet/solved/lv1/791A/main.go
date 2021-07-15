package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	count := 0

	for a <= b {
		count++
		a = a * 3
		b = b * 2
	}
	fmt.Println(count)
}
