package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a)
	for sumdigit(a)%4 != 0 {
		a++
	}
	fmt.Println(a)
}
func sumdigit(a int) int {
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	return sum
}
