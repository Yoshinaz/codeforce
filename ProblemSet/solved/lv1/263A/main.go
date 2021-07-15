package main

import (
	"fmt"
)

func main() {
	var ai, aj, a int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scanf("%d", &a)
			//fmt.Printf("(%d %d %d)", i, j, a)
			if a == 1 {
				ai = i
				aj = j
			}
		}
		fmt.Scanf("\n")
	}
	fmt.Println(abs(ai-2) + abs(aj-2))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
