package main

import (
	"bufio"
	"fmt"
	"os"
)
var N int
var p[] int
var dp[301][301] int
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
func max(a int,b int) int{
	if(a>b){
		return a
	}else{
		return b
	}
}
func main(){
	defer writer.Flush()
	fmt.Fscan(reader,&N)
	for i:=0;i<N;i++{
		var point int
		fmt.Fscan(reader,&point)
		p=append(p,point)
	}
	if N==1{
		fmt.Fprint(writer,p[0])
		return
	}
	dp[0][0]=0
	dp[0][1]=p[0]
	dp[1][0]=p[1]+p[0]
	dp[1][1]=p[1]
	for i:=2;i<N;i++{
		dp[i][0]=dp[i-1][1]+p[i]
		dp[i][1]=max(dp[i-2][0],dp[i-2][1])+p[i]
	}
	fmt.Fprint(writer,max(dp[N-1][0],dp[N-1][1]))
}