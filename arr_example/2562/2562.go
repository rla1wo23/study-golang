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
	var a [9]int
	max:=0
	maxIdx:=0
	for i:=0; i<9; i++{
		var k int
		fmt.Fscan(reader,&k)
		a[i] = k
		if max<k{
			max=k
			maxIdx=i
		}
	}
	fmt.Fprintf(writer, "%d \n%d",max,maxIdx+1)
}