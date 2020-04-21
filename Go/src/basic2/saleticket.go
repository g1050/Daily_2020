package main

import ("fmt"
        "sync"
)

var ticket = 100
var mutex sync.Mutex //创建一把锁
var wg sync.WaitGroup //同步等待组

func sale(name string){
    for{
        mutex.Lock()
        if ticket > 0{
            fmt.Println(name," sales ",ticket)
            ticket--
        }else{
            mutex.Unlock()
            fmt.Println("Ticket is over!")
            break;
        }
        mutex.Unlock()
    }
    wg.Done()
}

func main(){

    wg.Add(4)
    go sale("one ")
    go sale("two ")
    go sale("three ")
    go sale("four ")

    wg.Wait()
    fmt.Println("Hello, World!")
}

