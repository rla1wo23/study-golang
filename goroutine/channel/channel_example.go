package main

import (
	"fmt"
	"time"
)
func pinger(c chan string) {
    for i := 0; i < 5; i++ {
        fmt.Println("Sending ping...")
        c <- "ping" // 값을 보내고, 받는 쪽이 받을 때까지 대기. 즉, 채널 변수는 버퍼를 상대방이 비울 때 까지 기다린다.
    }
}
func printer(c chan string) {
    for i := 0; i < 5; i++ {
        msg := <-c // 값을 받을 때까지 대기. c 버퍼에 뭐가 들어가기 전까지 기다림. lock과 signal 구조가 추상화 돼 있다.
        fmt.Println("Received:", msg)
        time.Sleep(time.Second * 1) // 1초 대기
    }
}

func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go printer(c)

    time.Sleep(time.Second * 6) // 메인 고루틴 대기
}
