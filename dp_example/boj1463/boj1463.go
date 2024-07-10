package main

import (
	"bufio"
	"fmt"
	"os"
)
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
func min(a int,b int) int{
	if(a>b){
		return b
	}
	return a
}
var d[1000001] int
func dp(n int) int{
	if d[n]==10000005{
		if n%3==0&&n%2==0{
			d[n]=min(min(dp(n/3),dp(n/2)),dp(n-1))+1
		}else if n%3==0{
			d[n]=min(dp(n/3),dp(n-1))+1
		}else if n%2==0{
			d[n]=min(dp(n/2),dp(n-1))+1
		}else{
			d[n]=dp(n-1)+1
		}
	}
	return d[n]
}
func main(){
	defer writer.Flush()
	var N int
	fmt.Fscan(reader,&N)
	for i:=0;i<=N;i++{
		d[i]=10000005
	}
	d[1]=0
	fmt.Fprint(writer,dp(N))
}