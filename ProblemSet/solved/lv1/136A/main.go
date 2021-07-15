package main

import "fmt"

func main() {
	var n, a int
	fmt.Scanf("%d\n", &n)
	out := [100]int{}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		out[a-1] = i + 1
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", out[i])
	}
}
