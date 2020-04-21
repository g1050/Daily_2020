package main

import "fmt"
import "time"

func test(ch chan int){
    for value := range ch {
        fmt.Println("读到的数据是:",value)
    }
}

func main(){
    ch := make(chan int,6)
    defer close(ch)

    go test(ch)
    for i := 0;i<100;i++{
        ch <- i
        fmt.Println("\t写入的数据是:",i)
    }
    
    time.Sleep(3*time.Second)
    fmt.Println("Hello, World!")
}

