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

	var n int
	fmt.Fscanln(reader,&n)

	for i:=1; i<=n ; i++{
		for j:=i+1; j<=n ;j++{
			fmt.Fprint(writer," ")
		}
		for j:=n-i;j<n;j++{
			fmt.Fprint(writer,"*")
		}
		fmt.Fprint(writer,"\n")
	}
}