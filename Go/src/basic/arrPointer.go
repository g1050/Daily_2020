package main

import "fmt"

func main(){
    arr := [4]int{}
    var p *[4]int = &arr
    fmt.Println(arr)
    fmt.Println(p)
    (*p)[0] = 100
    fmt.Println(arr)
    p[0] = 200
    fmt.Println(arr)
}


