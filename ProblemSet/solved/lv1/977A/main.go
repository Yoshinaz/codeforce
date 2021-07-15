package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	for i := 0; i < k; i++ {
		if n%10 == 0 {
			n = n / 10
		} else {
			n = n - 1
		}
	}
	fmt.Println(n)
}
