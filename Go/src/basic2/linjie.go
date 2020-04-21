package main

import "fmt"
import "time"

func fun(p *int){
    *p = 2
    fmt.Println("fun a->",*p)
}

func main(){
    a := 1

    go fun(&a)
    time.Sleep(1*time.Second)
    fmt.Println("main a->",a)
    fmt.Println("Hello, World!")
}

