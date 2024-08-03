package main

import (
	"fmt"
	"sync"
)

// 전역 변수 선언
var a, b, c int64

// 재귀 함수 정의
func funcRecursive(exp int64) int64 {
    if exp == 1 {
        return a % c
    }
    x := funcRecursive(exp / 2) % c
    if exp % 2 == 0 {
        return x * x % c
    } else {
        return (x * x % c) * a % c
    }
}

// worker goroutine to perform multiplication and modulo
func worker(a, c, start, end int64, ch chan int64, wg *sync.WaitGroup) {
    defer wg.Done()
    result := int64(1)
    for i := start; i < end; i++ {
        result *= a
        result %= c
    }
    ch <- result
}

func main() {
    fmt.Scan(&a, &b, &c)

    numGoroutines := 8 // Number of goroutines to use
    chunkSize := b / int64(numGoroutines)

    var wg sync.WaitGroup
    ch := make(chan int64, numGoroutines)

    // Start goroutines
    for i := int64(0); i < int64(numGoroutines); i++ {
        start := i * chunkSize
        end := start + chunkSize
        if i == int64(numGoroutines)-1 {
            end = b
        }
        wg.Add(1)
        go worker(a, c, start, end, ch, &wg)
    }

    // Wait for all goroutines to finish
    wg.Wait()
    close(ch)

    // Combine results
    k := int64(1)
    for result := range ch {
        k *= result
        k %= c
    }

    fmt.Println(k)
}
