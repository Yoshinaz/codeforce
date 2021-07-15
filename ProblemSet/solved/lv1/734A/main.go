package main

import "fmt"

func main() {
	var n int
	var input string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &input)
	countA := 0
	countD := 0
	for _, v := range input {
		if v == 'A' {
			countA++
		} else {
			countD++
		}
	}
	if countA > countD {
		fmt.Println("Anton")
	} else if countA < countD {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}
