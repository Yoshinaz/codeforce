package main

import "fmt"

func main() {
	var n, k, score, kscore int
	fmt.Scanf("%d %d\n", &n, &k)
	pass := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &score)
		if score == 0 {
			break
		}
		if i+1 == k {
			kscore = score
		}
		if i+1 > k && score < kscore {
			break
		}
		pass++
	}
	fmt.Println(pass)
}
