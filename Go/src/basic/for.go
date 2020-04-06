package main

import "fmt"

func main(){
    for i := 58;i>=23;i--{
        fmt.Println(i)
    }

    sum := 0
    for i := 1;i<=100;i++{
        sum += i
    }
        fmt.Println(sum)
}

