package main

import "fmt"

func main(){
    increment := fun() //此时匿名函数并不执行,因为加()才会执行
    res := increment()
    fmt.Println(res)
    res = increment()
    fmt.Println(res)
}

func fun() func()int{
    i := 0
    increment := func ()int{
        i++
        return i
    }
    return increment
}

