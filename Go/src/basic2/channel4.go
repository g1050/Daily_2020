package main

import "fmt"
import "time"

func test(ch chan int){
    for value := range ch{
        fmt.Println("Value is ",value)    
    }
}

func main(){
    ch := make(chan int)
    defer close(ch)
    go test(ch)
    for i := 0;i<10;i++{
        ch <- i
    }
    time.Sleep(time.Second*3)
}

