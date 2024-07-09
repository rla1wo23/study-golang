package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	writer:=bufio.NewWriter(os.Stdout)
	var a,b,c int
	fmt.Fscanln(reader,&a,&b,&c)
	defer writer.Flush()
	ans:=0
	if (a==b)&&(b==c){ //셋 다 같은 경우
		ans=10000+a*1000
	}else if a==b { //하나만 같은 경우
		ans=1000+a*100	
	}else if a==c {
		ans=1000+a*100
	}else if b==c{
		ans=1000+b*100
	}else{
		if a>b&&a>c{
			ans=100*a
		}else if b>c{
			ans=100*b
		}else{
			ans=100*c
		}
	}
	fmt.Fprintln(writer,ans)
}