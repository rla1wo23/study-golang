package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var str string
	var N int
	fmt.Fscan(reader,&N)
	fmt.Fscan(reader,&str)
	ans:=0
	for i:=0;i<N;i++{
		ans+=int(str[i]-'0')
	}
	fmt.Fprint(writer,ans)
}