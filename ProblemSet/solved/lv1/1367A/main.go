package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var t int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&t)
	for i:=0;i<t;i++{
		fmt.Fscan(in,&s)
		fmt.Printf("%s",string(s[0]))
		for j:=1;j<len(s)-1;j+=2{
			fmt.Printf("%s",string(s[j]))
		}
		fmt.Printf("%s\n",string(s[len(s)-1]))
	}
}