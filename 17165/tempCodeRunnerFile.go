package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader* bufio.Reader=bufio.NewReader(os.Stdin)
var writer* bufio.Writer=bufio.NewWriter(os.Stdout)
func min(a,b int) int{
	if(a>b){
		return b
	}
	return a
}
func main(){
	defer writer.Flush()
	var N int
	var str string
	fmt.Fscan(reader,&N)
	fmt.Fscan(reader,&str)
	s:=[]rune(str)
	var r_s,b_s,r_e,b_e int
	var rc,bc int
	if s[0]=='R'{
		rc++
	}else{
		bc++
	}
	i:=1
	for ;s[i]==s[i-1]&&i<N;i++{
		if s[i]=='R'{
			r_s++
			rc++
		}else{
			b_s++
			bc++
		}
	}
	//이제 i는 다른게 됐음.
	j:=len(s)-1
	if s[j]=='R'{
		rc++
		r_e++
	}else{
		bc++
		b_e++
	}
	j--
	for ;s[j]==s[j+1]&&j>=0;j--{
		if s[j]=='R'{
			r_e++
			rc++
		}else{
			b_e++
			bc++
		}
	}
	for ;i<=j&&i<N;i++{
		if s[i]=='R'{
			rc++
		}else{
			bc++
		}
	}
	ans:=min(rc-r_e,bc-b_s)
	ans=min(ans,bc-b_e)
	ans=min(ans,rc-r_s)
	fmt.Fprintln(writer,ans)
}