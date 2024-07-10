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
	var N int
	fmt.Fscan(reader,&N)
	var arr []int
	for i:=0;i<N;i++{
		var k int
		fmt.Fscan(reader,&k)
		arr=append(arr,k)
	}
	for i:=0;i<N-1;i++{
		for j:=0;j<N-1-i;j++{
			if arr[j]>arr[j+1]{
				tmp:=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=tmp
			}
		}
	}
	for i:=0;i<N;i++{
		fmt.Fprintln(writer,arr[i])
	}
}