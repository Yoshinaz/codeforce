package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	var line string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	result := make([]string, 0)
	hasPair := false
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &line)
		if !hasPair {
			s := strings.Split(line, "|")
			if s[0] == "OO" {
				hasPair = true
				result = append(result, "++|"+s[1])
			} else if s[1] == "OO" {
				hasPair = true
				result = append(result, s[0]+"|++")
			} else {
				result = append(result, line)
			}
		} else {
			result = append(result, line)
		}
	}
	if hasPair {
		fmt.Println("YES")
		for i := 0; i < n; i++ {
			fmt.Println(result[i])
		}
	} else {
		fmt.Println("NO")
	}
}
