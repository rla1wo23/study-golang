package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var reader* bufio.Reader=bufio.NewReader(os.Stdin)
	var str string
	str, _= reader.ReadString('\n')
	fmt.Println(str)
}