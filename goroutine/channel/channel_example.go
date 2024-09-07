package main

import (
	"fmt"
	"time"
)
func pinger(c chan string) {
    for i := 0; i < 5; i++ {
        fmt.Println("Sending ping...")
        c <- "ping" // 값을 보내고, 받는 쪽이 받을 때까지 대기
    }
}
func printer(c chan string) {
    for i := 0; i < 5; i++ {
        msg := <-c // 값을 받을 때까지 대기
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
