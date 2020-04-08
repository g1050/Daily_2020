package main

import "fmt"

func main(){
    a := 1 
    b := 2
    c := 3
    d := 4

    fmt.Println(a,b,c,d)
    var arr [4]*int //指针数组

    arr[0] = &a
    arr[1] = &b
    arr[2] = &c
    arr[3] = &d

    *arr[0] = 100
    fmt.Println(a,b,c,d)
}




