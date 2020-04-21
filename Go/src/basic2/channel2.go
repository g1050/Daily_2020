package main

import "fmt"

func test(ch chan bool){
    for i:=0;i<1000;i++{
        fmt.Println(i)
    }
    ch <- true
}

func main(){
    ch  := make(chan bool)
    go test(ch)

    data := <- ch //从通道中读取数据
    if data {
        fmt.Println("Over!")
    }
}

