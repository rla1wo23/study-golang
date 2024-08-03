package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader* bufio.Reader=bufio.NewReader(os.Stdin)
var writer* bufio.Writer=bufio.NewWriter(os.Stdout)
var pc[] int
func main(){
	var N,M int
	fmt.Fscan(reader,&N)
	fmt.Fscan(reader,&M)
	for i:=0;i<N;i++{
		pc=append(pc,i)
	}
	for i:=0;i<M;i++{

	}
	
	
}