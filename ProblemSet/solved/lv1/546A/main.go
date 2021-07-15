package main

import "fmt"

func main() {
	var k, n, w int
	fmt.Scanf("%d %d %d\n", &k, &n, &w)
	price := w * (w + 1) / 2
	price *= k
	if n >= price {
		fmt.Println("0")
	} else {
		fmt.Println(price - n)
	}

}
