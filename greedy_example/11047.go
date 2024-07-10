package main

import (
	"bufio"
	"fmt"
	"os"
)
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
func main(){
	defer writer.Flush()
	var N, K int
	fmt.Fscanln(reader,&N, &K)
	var arr []int
	for i:=0;i<N;i++{
		num:=0
		fmt.Fscan(reader,&num)
		arr=append(arr,num)
	}
	arr=append(arr,100000001)
	coins:=0
	for K!=0{
		maxIdx:=0
		for arr[maxIdx+1]<=K{
			maxIdx++
		}	
		factor:=K/arr[maxIdx]
		coins+=factor
		K=K%arr[maxIdx]
	}
	fmt.Fprint(writer,coins)
}