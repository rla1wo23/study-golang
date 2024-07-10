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
	var N int
	var nums []int
	fmt.Fscan(reader,&N)
	for i:=0; i<N;i++{
		var k int
		fmt.Fscan(reader,&k)
		nums=append(nums,k)
	}
	var target int
	cnt:=0
	fmt.Fscan(reader,&target)
	
	for i:=0; i<N;i++{
		if nums[i]==target{
			cnt++
		}
	}
	fmt.Fprint(writer,cnt)
}