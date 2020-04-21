package main

import "fmt"
import "time"

func main(){
    ch1 := make(chan int)
    ch2 := make(chan int)
    defer close(ch1)
    defer close(ch2)

    go func(){
        time.Sleep(time.Second*3)
        ch1 <- 100
    }()

    go func(){
        time.Sleep(time.Second*3)
        ch2 <- 200
    }()

    select {
        case data := <- ch2:{
            fmt.Println(data)
        }
        case data,ok := <- ch1:{
            if ok{
                fmt.Println("ch1读取到数据",data)
            }else{
                fmt.Println("ch1没有读取到数据")
            }
        }
        case <- time.After(3*time.Second):{
            fmt.Println("定时器")
        }

    // default:
    //     fmt.Println("default")

    }
}

