package main

import (
	"bufio"
	"fmt"
	"os"
)
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
func main(){
	defer writer.Flush()
	var N,K int
	fmt.Fscanln(reader,&N,&K)
	var days[] int
	for i:=0;i<N;i++{
		var deg int
		fmt.Fscan(reader,&deg)
		days=append(days,deg)
	}
	initial_sum:=0
	var rec[] int
	for i:=0;i<K-1;i++{
		rec=append(rec,-101)
		initial_sum+=days[i]
	}
	initial_sum+=days[K-1]
	rec=append(rec,initial_sum)
	max:=initial_sum
	for i:=K;i<N;i++{
		today:=rec[i-1]+days[i]-days[i-K]
		if(max<today){
			max=today
		}
		rec=append(rec,today)
	}
	fmt.Fprint(writer,max)
}