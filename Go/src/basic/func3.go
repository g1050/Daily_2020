package main

import "fmt"

func main(){
    fmt.Printf("%d +  %d = %d\n",1,2,operate(1,2,add))
}

func add(a,b int)int{
    return a+b
}

func operate(a,b int ,fun func(int,int)int)int{
    res := fun(a,b)
    return res
}

