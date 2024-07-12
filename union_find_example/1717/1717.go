package main

import (
	"bufio"
	"fmt"
	"os"
)
var n,m int
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
var arr[] int
func find(key int) int{ //근본 찾기
	if(key==arr[key]){
		return key
	}else{
		arr[key]=find(arr[key])
		return arr[key]
	}
}
func union(a,b int){
	arr[find(arr[b])]=find(arr[a])
}
func is_union(a,b int){
	if find(a)==find(b){
		fmt.Fprintln(writer,"YES")
	}else{
		fmt.Fprintln(writer,"NO")
	}
}
func main(){
	defer writer.Flush()
	fmt.Fscanln(reader,&n,&m)
	for i:=0;i<=n;i++{
		arr=append(arr,i)
	}
	for i:=0;i<m;i++{
		var k,a,b int
		fmt.Fscanln(reader,&k,&a,&b)
		if k==0{
			union(a,b)
		}else{
			is_union(a,b)
		}
	}
}