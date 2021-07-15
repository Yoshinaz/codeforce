package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	in := bufio.NewReader(os.Stdin)
	s, _ = in.ReadString('\n')
	//fmt.Fscanln(in, &s)
	//fmt.Println(s)
	ss := s[1:strings.Index(s, "}")]
	if len(ss) == 0 {
		fmt.Println(0)
	} else {
		split := strings.Split(ss, ",")
		distince := make(map[string]bool, 0)
		for _, v := range split {
			distince[strings.TrimSpace(v)] = true
		}
		ans := 0
		for _, _ = range distince {
			ans++
		}
		fmt.Println(ans)
	}
}
