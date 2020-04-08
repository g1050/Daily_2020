package main

import "fmt"


func main(){

    var num = 1

    Label:
    for num<= 20{
        if num == 15{
            num++
            goto Label
        }
        fmt.Println(num)
        num++
    }
}

