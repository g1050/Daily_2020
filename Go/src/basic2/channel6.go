package main

import "fmt"

func test(ch <- chan int,ch2 chan bool){
    for value := range ch{
        fmt.Println(value)
    }
    ch2 <- true
}

func main(){
    ch := make(chan int)//创建一个双向通道
    defer close(ch)
    ch2 := make(chan bool)
    go test(ch,ch2)
    for i := 0;i<10;i++{
        ch <- i
    }
    data := <- ch2
    fmt.Println("Hello, World!",data)
}


