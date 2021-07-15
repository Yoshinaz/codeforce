package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(math.Ceil(float64(n) / 5))
}
