package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var n, h int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&n)
	min := 101
	max := 0
	minIdx := 0
	maxIdx := 0
	for i:=0;i<n;i++{
		fmt.Fscan(in,&h)
		if h<=min{
			min = h
			minIdx = i
		}
		if h>max{
			max = h
			maxIdx = i
		}
	}
	if maxIdx>minIdx{
		fmt.Println(maxIdx + n-minIdx-2)
	}else{
		fmt.Println(maxIdx + n-minIdx-1)
	}
}