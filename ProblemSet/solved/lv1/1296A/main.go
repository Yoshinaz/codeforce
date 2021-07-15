package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var t,n,a int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&t)
	for i:=0;i<t;i++{
		sum :=0
		hasEven := false
		hasOdd := false
		fmt.Fscan(in,&n)
		for j:=0;j<n;j++{
			fmt.Fscan(in,&a)
			sum +=a
			if a%2==0{
				hasEven = true
			}else{
				hasOdd = true
			}
		}
		if sum%2 ==1{
			fmt.Println("YES")
		}else{
			if hasEven && hasOdd{
				fmt.Println("YES")
			}else{
				fmt.Println("NO")
			}
		}
	}
}