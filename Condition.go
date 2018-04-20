package main

import (
    "fmt"
    "sync"
    "time"
    "container/list"
)

var x int = 0
var cond = sync.NewCond(&sync.Mutex{})
var queue = list.New()

func producer(num int){
    for {
        cond.L.Lock()
        for queue.Len() == 100 {
            cond.Wait()
        }
        queue.PushBack(num)
        cond.Signal()
        cond.L.Unlock()
    }
}

func consumer(num int){
    for {
        cond.L.Lock()
        for queue.Len() == 0 {
            cond.Wait()
        }
        fmt.Printf("Qsize %d, consumer %d && worker %d\n", queue.Len(), num, queue.Front().Value)
        queue.Remove(queue.Front())
        cond.Signal()
        cond.L.Unlock()
    }
}

func main(){
    for i := 0 ; i < 10; i++ {
        go producer(i)
    }
    for i := 0 ; i < 2; i++ {
        go consumer(i)
    }
    time.Sleep(time.Second * 5)
}
