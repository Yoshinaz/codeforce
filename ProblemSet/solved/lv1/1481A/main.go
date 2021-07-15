package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, x, y int
	var s string
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		data := make([]int, 4)
		fmt.Fscan(in, &x, &y, &s)
		for _, v := range s {
			data[getIdx(v)]++
		}
		if abs(x) <= data[getXY(x)] && abs(y) <= data[getUD(y)] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getUD(y int) int {
	if y > 0 {
		return 0
	}
	return 2
}
func getXY(x int) int {
	if x > 0 {
		return 1
	}
	return 3
}
func getIdx(r rune) int {
	switch r {
	case 'U':
		return 0
	case 'R':
		return 1
	case 'D':
		return 2
	case 'L':
		return 3
	}
	return 0
}
