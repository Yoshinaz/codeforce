package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, call int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tmp := []int{1, 3, 6, 10}
	fmt.Sscanf(scanner.Text(), "%d", &n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		room := scanner.Text()
		fmt.Sscanf(room, "%d", &call)
		ans := (call%10-1)*10 + tmp[len(room)-1]
		fmt.Println(ans)
	}
}
