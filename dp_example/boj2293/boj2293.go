package main

import (
	"bufio"
	"fmt"
	"os"
)
var n,k int
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)

var coins [] int
var dp[100001] int
func main(){
	defer writer.Flush()
	fmt.Fscanln(reader,&n,&k)
	for i:=0;i<n;i++{
		var num int
		fmt.Fscan(reader,&num)
		coins=append(coins,num)
	}
	dp[0]=1
	for i:=0;i<n;i++{
		c:=coins[i]
		for j:=c;j<=k;j++{
			dp[j]+=dp[j-c]
		}
	}
	fmt.Fprintln(writer,dp[k])
}