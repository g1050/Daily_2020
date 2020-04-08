package main

import "fmt"

func main(){

    fmt.Println(getSum(1,2))
    
    function1 := getSum

    fmt.Println(getSum,function1)

    fmt.Println(function1(1,2))

    var function2 func(int,int)(int)
    function2 = getSum
    fmt.Println(function2)
}

func getSum(a,b int)int{
    return a+b
}

