package main

import (
	"bufio"
	"fmt"
	"os"
)
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
var n,m int
var pred []int
var pwd [7]int
var isused[10] int
var cnt = 0
func brute(length int){
	if(length==n){
		for i:=0;i<len(pred);i++{
			if isused[pred[i]]==0{
				return
			}
		}
		cnt++
		return
	}
	for i:=0;i<10;i++{
		isused[i]++
		pwd[length]=i
		brute(length+1)
		isused[i]--
	}
}

func main(){
	defer writer.Flush()
	fmt.Fscanln(reader,&n,&m)
	for i:=0;i<m;i++{
		var k int
		fmt.Fscan(reader,&k)
		pred=append(pred,k)
	}
	brute(0)
	fmt.Fprint(writer,cnt)
}