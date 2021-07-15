package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&n)
	a := make([]int,n)
	for i:=0;i<n;i++{
		fmt.Fscan(in,&a[i])
	}
	sort.Ints(a)
	sum := sum(a)
	half := sum/2+1
	
	sum =0
	count := 0
	for i:= n-1;i>=0;i--{
		sum+=a[i]
		count++
		if sum>=half{
			break
		}
	}
	fmt.Println(count)
}

func sum(a []int)int{
	sum :=0
	for _,v := range a{
		sum+=v
	}
	return sum
}