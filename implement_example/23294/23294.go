package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var N, Q, C int
var capacity, stack_size int
var now int = 0
var pages []int

type deque []int
type stack []int

var q deque
var s stack

func (q *deque) push_back(i int) {
	(*q) = append(*q, i)
}

func (q *deque) pop_front() int {
	item := (*q)[0]
	(*q) = (*q)[1:]
	return item
}

func (q *deque) pop_back() int {
	item := (*q)[len(*q)-1]
	(*q) = (*q)[:len(*q)-1]
	return item
}

func (s *stack) pop() int {
	item := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return item
}

func (s *stack) push(item int) {
	(*s) = append(*s, item)
}

func (s *stack) clear() {
	(*s) = (*s)[:0]
}

func Access(p int) {
	s.clear()
	capacity -= stack_size
	stack_size = 0
	if now != 0 {
		q.push_back(now)
	}
	now = p
	capacity += pages[p]

	for capacity > C {
		item := q.pop_front()
		capacity -= pages[item]
	}
}

func Backward() {
	if len(q) == 0 {
		return
	}
	s.push(now)
	stack_size += pages[now]
	now = q.pop_back()
}

func Forward() {
	if len(s) == 0 {
		return
	}
	q.push_back(now)
	now = s.pop()
	stack_size -= pages[now]
}

func Compress() {
	for i := len(q) - 1; i > 0; i-- {
		if q[i] == q[i-1] {
			item := q[i]
			capacity -= pages[item]
			q = append(q[:i], q[i+1:]...)
		}
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscanln(reader, &N, &Q, &C)
	pages = append(pages, 0)
	for i := 1; i <= N; i++ {
		var k int
		fmt.Fscan(reader, &k)
		pages = append(pages, k)
	}
	for i := 0; i < Q; i++ {
		var cmd string
		fmt.Fscan(reader,&cmd)
		if cmd=="A"{
			var k int
			fmt.Fscan(reader,&k)
			Access(k)
		} else if cmd == "B" {
			Backward()
		} else if cmd == "F" {
			Forward()
		} else if cmd == "C" {
			Compress()
		}
	}
	fmt.Fprintln(writer, now)
	if len(q) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		for i := len(q) - 1; i >= 0; i-- {
			fmt.Fprintf(writer, "%d ", q[i])
		}
		fmt.Fprintln(writer)
	}
	if len(s) == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		for len(s) != 0 {
			fmt.Fprintf(writer, "%d ", s.pop())
		}
	}
}
