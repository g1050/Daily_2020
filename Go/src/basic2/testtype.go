package main

import "fmt"
type myint int

func main(){
    var num myint = 100
    var num2 int = 200
    num = (int)num2
    fmt.Println(num)
}

