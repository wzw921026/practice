package main

import (
    "fmt"
    "time"
)

var _chan = make(chan int)

func producer(num int){
    for {
        _chan <- num
        time.Sleep(time.Second / 10)
    }
}

func consumer(num int){
    var x int
    var ok bool
    for {
        x ,ok = <-_chan
        if ok {
            fmt.Printf("consumer %d && producer %d\n", num, x)
        } else {
            fmt.Printf("NULL")
        }
    }
}

func main(){
    for i := 0; i < 2; i++ {
        go producer(i)
    }
    for i := 0; i < 10; i++ {
        go consumer(i)
    }
    time.Sleep(time.Second * 10)
}
