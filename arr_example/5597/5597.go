package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	writer:=bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var arr[31] int
	arr[0]=1
	for i:=0;i<28;i++{
		var k int
		fmt.Fscan(reader,&k)
		arr[k]=1
	}
	for i:=0;i<31;i++{
		if arr[i]==0{
			fmt.Fprintln(writer,i)
		}
	}
}