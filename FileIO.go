package main

import (
    "os"
    "io"
)

func main() {
    f, _ := os.OpenFile("output", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    defer f.Close()
    io.WriteString(f, "hello world\n")
}
