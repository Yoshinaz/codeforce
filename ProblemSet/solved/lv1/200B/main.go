package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var n int
	var p float64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&n)
	sum := 0.0
	for i:=0;i<n;i++{
		fmt.Fscan(in,&p)
		sum +=p
	}
	fmt.Printf("%0.6f\n",sum/float64(n))
}