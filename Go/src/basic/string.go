package main

import "fmt"

func main(){
    name := "Wanghuahua"
    for index,value := range name{
        fmt.Printf("%d %c\n",index,value)
    }
}

