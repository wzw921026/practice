package main

import (
    "fmt"
    "sync"
)

var (
    x int
    mutex sync.Mutex
    wait_group sync.WaitGroup
)

func calc(num int){
    for i := 0; i < 100000; i++ {
        mutex.Lock()
        x++
        mutex.Unlock()
    }
    fmt.Printf("th %d ok\n", num);
    wait_group.Done()
}

func main(){
    var n = 10
    wait_group.Add(n)
    for i := 0 ; i < n; i++ {
        go calc(i)
    }
    wait_group.Wait()
    fmt.Printf("x=%d\n", x)
}
