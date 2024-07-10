package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
var N,M int

var arr []int
func binary_search(target int) int{
	st:=0
	en:=N-1
	mid:=0
	for st<=en {
		mid=(st+en)/2
		if(target==arr[mid]){
			return 1
		}else if(target>arr[mid]){
			st=mid+1
		}else{
			en=mid-1
		}
	}
	return 0
}

func main(){
	defer writer.Flush()
	fmt.Fscan(reader,&N)
	for i:=0;i<N;i++{
		var k int
		fmt.Fscan(reader,&k)
		arr=append(arr,k)
	}
	fmt.Fscan(reader,&M)
	sort.Ints(arr)
	for i:=0;i<M;i++{
		var t int
		fmt.Fscan(reader,&t)
		fmt.Fprintln(writer,binary_search(t))
	}
}