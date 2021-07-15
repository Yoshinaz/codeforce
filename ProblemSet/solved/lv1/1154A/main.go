package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	var input [4]int
	fmt.Fscan(in,&input[0],&input[1],&input[2],&input[3])
	max,left := max4(input)
	//fmt.Println(max,left)
	for _,v := range left{
		fmt.Printf("%d ",max -v )
	}
}

func max4(input [4]int) (int,[]int){
	maxIdx := 0
	for i:=0;i<len(input);i++{
		if input[maxIdx]<input[i]{
			maxIdx = i
		}
	}
	//fmt.Println(maxIdx)
	max := input[maxIdx]
	return max,append(input[0:maxIdx],input[maxIdx+1:len(input)]...)
}