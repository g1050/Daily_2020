package main

import "fmt"
import "time"

func test(ch chan int){
    for{
        // time.Sleep(1*time.Second)
        data,ok := <- ch
        if ok{
            fmt.Println("读到数据",data,ok)
        }else{
            fmt.Println("数据已经读取完毕",data,ok)
            break;
        }
    }
}

func main(){
    
    ch := make(chan int)
    go test(ch)

    for i:= 0;i<10;i++{
        ch <- i
    }
    close(ch)
    time.Sleep(3*time.Second)
    fmt.Println("Hello, World!")
}

