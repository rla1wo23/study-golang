package main

import (
	"bufio"
	"fmt"
	"os"
)
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
var dp[1000001] int

func main(){
	defer writer.Flush()
	var N int
	fmt.Fscan(reader,&N)
	
	for N!=0{
		var str string
		str, _=reader.ReadString('\n')
		s:=[]rune(str)
		for i:=0;i<len(s);i++{
			dp[i]=0
		}
		var m map[rune]int
		m=make(map[rune]int)
		m[s[len(s)-1]]=1
		cnt:=1//초기 길이
		max:=1
		i:=len(s)-2
		dp[len(s)-1]=1
		for ;i>0&&cnt<=N;i--{
			if m[s[i]]==1{
				dp[i-1]=dp[i]+1
				continue
			}else{
				m[s[i]]=1
				cnt++
				if cnt<=N{
					dp[i-1]=dp[i]+1
				}
			}
		}
		for ;i>=0;i--{
			forwarding:=0
			j:=i
			for ;s[j]!=s[i]&&j<len(s)-1;j++{
				forwarding++
			}
			forwarding+=dp[j]
			for ;s[j]!=s[i]&&j<len(s)-1;j++{
				forwarding++
			}
			dp[j]=forwarding
			if dp[j]>max{
				max=dp[j]
			}
		}
		fmt.Fprint(writer,max)
		fmt.Fscan(reader,&N)
	}
}