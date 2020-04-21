package main

import "fmt"
import "time"

func main(){
    timer := time.NewTimer(3*time.Second)

    fmt.Printf("%T\n",timer)
    fmt.Println(time.Now())

    ch2 := timer.C
    fmt.Println(<-ch2)
    
    fmt.Println("Hello, World!")
}

