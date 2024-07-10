package main

import (
	"bufio"
	"fmt"
	"os"
)
var dp[12] int
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
func fnd(target int) int{
	if dp[target]!=0{
		return dp[target]
	}
	return fnd(target-1)+fnd(target-2)+fnd(target-3)
}
func main(){
	defer writer.Flush()
	var N int
	fmt.Fscan(reader,&N)
	dp[0]=-1
	dp[1]=1
	dp[2]=2
	dp[3]=4
	for i:=0;i<N;i++{
		var target int
		fmt.Fscan(reader,&target)
		fmt.Fprintln(writer,fnd(target))
	}
}