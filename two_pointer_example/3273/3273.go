package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
var arr[] int
func main(){
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader,&n)
	for i:=0;i<n;i++{
		var k int
		fmt.Fscan(reader,&k)
		arr=append(arr,k)
	}
	ans:=0
	var x int
	fmt.Fscan(reader, &x)
	sort.Ints(arr)
	st:=0
	en:=n-1
	for en>st {
		if (arr[st]+arr[en])==x{
			ans++
			en--
		}else if (arr[st]+arr[en])>x{
			en--
		}else if (arr[st]+arr[en])<x{
			st++
		}
	}
	fmt.Fprintln(writer,ans)
}