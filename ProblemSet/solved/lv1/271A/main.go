package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d\n", &input)
	for !isDistinct(input + 1) {
		input++
	}
	fmt.Println(input + 1)
}

func isDistinct(input int) bool {
	flag := [10]bool{}
	for input > 0 {
		if flag[input%10] {
			return false
		}
		flag[input%10] = true
		input /= 10
	}
	return true
}
