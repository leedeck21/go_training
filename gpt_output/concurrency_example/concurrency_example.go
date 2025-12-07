package main

import (
	"fmt"
	"time"
)

func fetchData(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "Data loaded"
}

func main() {
    ch := make(chan string)
    go fetchData(ch) // non-blocking
    fmt.Println("Waiting for data...")
    msg := <-ch // block until goroutine sends
    fmt.Println(msg)
}
