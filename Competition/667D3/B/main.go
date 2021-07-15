package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t, n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &n)
		scanner.Scan()
		books := strings.Split(scanner.Text(), " ")
		lbook := -1
		rbook := 0
		nbook := 0
		for i, v := range books {
			if lbook == -1 && v == "1" {
				lbook = i
				rbook = i
			} else if lbook != -1 && v == "1" {
				rbook = i
				nbook++
			}
		}
		space := rbook - lbook - nbook
		//spaceL := lbook
		//spaceR := len(books) - rbook - 1
		ans := space //+ min(spaceL, spaceR)
		//fmt.Println(lbook, rbook, nbook)
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
