package main

import "fmt"

func main() {
	var r, c int
	fmt.Scanf("%d %d\n", &r, &c)
	ans := r / 2 * c
	if r%2 == 1 {
		ans += c / 2
	}
	fmt.Println(ans)
}
