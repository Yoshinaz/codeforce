package main

import "fmt"

func main() {
	var n, in, out int
	fmt.Scanf("%d\n", &n)
	count := 0
	max := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &out, &in)
		count = count - out + in
		if max < count {
			max = count
		}
	}
	fmt.Println(max)
}
