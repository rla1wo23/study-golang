package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	reader:=bufio.NewReader(os.Stdin)
	writer:=bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var str string
	fmt.Fscan(reader,&str)	
	length:=len(str)
	for i:=0; i<length/2;i++{
		if str[i]!=str[length-1-i]{
			fmt.Fprint(writer,"0")
			return
		}
	}
	fmt.Fprint(writer,"1")

}