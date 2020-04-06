package main

import "fmt"

func main(){
    var num int
    fmt.Scanf("%d",&num)
    if num > 10{
        fmt.Println(">10")
    }else{
        fmt.Println("<10")
    }

    if num2 := 10;num2 == 10{
        fmt.Println("num2 == 10")
    }

    // fmt.Println(num2)
}

