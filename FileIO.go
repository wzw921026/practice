package main

import (
    "os"
    "io"
)

func main() {
    f, err := os.OpenFile("output", os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil && os.IsNotExist(err) {
        f, _ = os.Create("output")
    }
    defer f.Close()
    io.WriteString(f, "hello world\n")
}
