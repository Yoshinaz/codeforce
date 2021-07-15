package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var table string
	var card [5]string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &table)
	for i := 0; i < 5; i++ {
		fmt.Fscan(in, &card[i])
	}
	ans := "NO"
	for i := 0; i < 5; i++ {
		if card[i][0] == table[0] || card[i][1] == table[1] {
			ans = "YES"
			break
		}
	}
	fmt.Println(ans)
}
