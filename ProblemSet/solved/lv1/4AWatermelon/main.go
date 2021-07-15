package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)
	mod := input % 2
	if mod == 0 && input > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
