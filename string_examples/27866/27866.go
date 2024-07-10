package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var reader *bufio.Reader=bufio.NewReader(os.Stdin)
	var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var str string
	fmt.Fscanln(reader,&str)
	var N int
	fmt.Fscan(reader,&N)
	fmt.Fprintf(writer,"%c",str[N-1])
}