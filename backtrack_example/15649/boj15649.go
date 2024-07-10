package main

import (
	"bufio"
	"fmt"
	"os"
)
var arr [8]int
var used [8]int
var lim int
var N int
var reader *bufio.Reader=bufio.NewReader(os.Stdin)

var writer *bufio.Writer=bufio.NewWriter(os.Stdout)

func tracker(k int){
		if k==lim{
			for i:=0;i<lim;i++{
				fmt.Fprintf(writer,"%d ",arr[i])
			}
			fmt.Fprintf(writer,"\n")
		}else{
			for i:=0;i<N;i++{
				if used[i]==0{
					used[i]=1
					arr[k]=i+1
					tracker(k+1)
					used[i]=0
				}
			}
		}
}
func main(){
	fmt.Fscan(reader,&N)	
	fmt.Fscan(reader,&lim)
	tracker(0)
	defer writer.Flush()
}