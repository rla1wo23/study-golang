package main

import (
	"bufio"
	"fmt"
	"os"
)
var N int
var reader *bufio.Reader=bufio.NewReader(os.Stdin)
var writer *bufio.Writer=bufio.NewWriter(os.Stdout)
type pair struct{
	First int
	Second int
}
type queue struct{
	items [] pair
}
func(q *queue) push(item pair){
	q.items=append(q.items,item)
}
func(q *queue) is_empty() bool{
	if len(q.items)==0{
		return true
	}else{
		return false
	}
}
func(q *queue) pop() pair{
	if len(q.items)==0{
		var ep pair
		ep.First=-1
		ep.Second=-1
		return ep
	}else{
		item:=q.items[0]
		q.items=q.items[1:]
		return item
	}
}
var board[100][100] int
var fire[100][100] int
var dx[3] int=[3]int{1,0,1}
var dy[3] int=[3]int{0,1,1}
func main(){
	fmt.Fscan(reader,&N)
	defer writer.Flush()
	var q queue 
	p1:=pair{First:1,Second:0}
	p2:=pair{First:1,Second:1}
	q.push(p1)
	q.push(p2)
	board[0][0]=-1
	board[1][0]=1
	board[1][1]=1
	ans:=0
	for !q.is_empty(){
		cur:=q.pop()
		X:=cur.First
		Y:=cur.Second
		if board[X][Y]==1||X==N-1{
				ans++
				continue
		}
		for i:=0;i<3;i++{
			nx:=X+dx[i]
			ny:=Y+dy[i]
			if nx>=N||ny>=N||nx<ny{
				continue
			}
			if board[X][Y]==1{//인간의 진행이라면
				if board[nx][ny]==-1 || board[nx][ny]==1{
					continue
				}

			}
			board[nx][ny]=board[X][Y]
			newp:=pair{First:nx,Second:ny}
			q.push(newp)
		}
	}
	fmt.Fprintln(writer,ans*2)
}