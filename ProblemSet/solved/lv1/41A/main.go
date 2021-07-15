package main

import "fmt"

func main() {
	var in1, in2 string
	fmt.Scanf("%s\n", &in1)
	fmt.Scanf("%s\n", &in2)
	if in1 == rev(in2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func rev(input string) string {
	if len(input) == 1 {
		return input
	} else {
		return input[len(input)-1:] + rev(input[0:len(input)-1])
	}
}
