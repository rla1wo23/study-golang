package main

import (
	"bufio"
	"fmt"
	"os"
)
var reader* bufio.Reader=bufio.NewReader(os.Stdin)
var writer* bufio.Writer=bufio.NewWriter(os.Stdout)
func main(){
	defer writer.Flush()
	var N int
	fmt.Fscan(reader,&N)
	var dp[102][102] int
	dp[2][1]=1
	for i:=3;i<=N;i++{
		for j:=1;j<i;j++{
			dp[i][j]=(dp[i-1][j]*2+dp[i-1][j+1]+dp[i-1][j-1])%10007
		}
	}
	ans:=0;
	for i:=1;i<=N;i++{
		ans+=dp[N][i]
	}
	fmt.Fprintln(writer,(ans*2)%10007)
}