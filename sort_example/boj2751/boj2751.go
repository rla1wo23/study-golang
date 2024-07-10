package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	var reader *bufio.Reader=bufio.NewReader(os.Stdin)
	var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var N int
	fmt.Fscan(reader,&N)
	var arr []int
	for i:=0;i<N;i++{
		var k int
		fmt.Fscan(reader,&k)
		arr=append(arr,k)
	}
	sort.Ints(arr)
	for i:=0;i<N;i++{
		fmt.Fprintln(writer,arr[i])
	}
}