package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var code string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &code)
	pos := 0
	ans := 0
	for _, v := range code {
		p := int(v - 'a')
		dis := abs(pos - p)
		dis = min(dis, 26-dis)
		ans += dis
		pos = p
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
