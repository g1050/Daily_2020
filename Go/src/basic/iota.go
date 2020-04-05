package main

import "fmt"

func main(){
    const(
        a = iota //0
        b = iota //1
        c = iota //2
    )

    
    fmt.Println(a,b,c)
}

