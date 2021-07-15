package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	var input string
	var out byte
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &input)
		fmt.Printf("1")
		sum := (input[0] - '0') + 1
		for j := 1; j < n; j++ {
			//fmt.Println("--", int(sum), string(input[j]))
			sum, out = getOut(input[j], sum)
			//fmt.Println(int(sum), string(out))
			fmt.Printf("%s", string(out))
		}
		fmt.Println()
	}
}

func getOut(a byte, prev byte) (byte, byte) {
	switch prev {
	case 2:
		if a == '1' {
			return 1, '0'
		}
		return 1, '1'
	case 1:
		if a == '1' {
			return 2, '1'
		}
		return 0, '0'
	case 0:
		return a - '0' + 1, '1'
	}
	return 1, '1'
}
