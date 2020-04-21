package main

import "fmt"
import "time"

func printNum(){
    for i:= 0;i<1000;i++{
        fmt.Println(i)
    }
}

func main(){



    for i:=0;i<1000;i++{
        fmt.Println("\ta")
    }
    go printNum()

    time.Sleep(1*time.Second)
    fmt.Println("Hello, World!")
}

