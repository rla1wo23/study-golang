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
	fmt.Fprintf(writer,"%d",len(str))
}