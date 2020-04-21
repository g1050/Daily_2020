package main

import "fmt"
import "sync"

// 创建同步等待组
var wg sync.WaitGroup
func main(){

    fmt.Println("Hello, World!")
    wg.Add(2)//计数器变为2

    go printNum()
    go printChar()

    fmt.Println("主goroutine阻塞")
    wg.Wait()
    fmt.Println("解除阻塞")

}

func printNum(){

    for i:=0;i<1000;i++{
        fmt.Println(i)
    }
    wg.Done()

}

func printChar(){

    for i:=0;i<1000;i++{
        fmt.Println("\ta")
    }
    wg.Done()

}
