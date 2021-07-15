package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var t,n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&t)
	for i:=0;i<t;i++{
		fmt.Fscan(in,&n)
		a := make([]int64, n)
		b := make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &b[j])
		}
		mina := min(a)
		minb := min(b)
		count := int64(0)
		for j := 0; j < n; j++ {
			da := a[j]-mina
			db := b[j]-minb
			count += max(da,db)
		}
		fmt.Println(count)
	}
	
}

func min(a []int64) int64{
	x := a[0]
	for i:=1;i<len(a);i++{
		if x>a[i]{
			x = a[i]
		}
	}
	return x
}
func max(a,b int64)int64{
	if a>b{
		return a
	}
	return b
}