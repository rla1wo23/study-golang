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
	var N int
	fmt.Fscan(reader,&N)
	var arr[] int
	for i:=0;i<N;i++{
		var k int
		fmt.Fscan(reader,&k)
		arr=append(arr,k)
	}
	var rec[1000] int
	for i:=N-2;i>=0;i--{
		for j:=i;j<N;j++{
			if rec[i]<rec[j]+1 && arr[i]<arr[j]{
				rec[i]=rec[j]+1
			}
		}
	}
	max:=0
	for i:=0;i<N;i++{
		if max<rec[i]{
			max=rec[i]
		}
	}
	fmt.Fprint(writer,max+1)
}