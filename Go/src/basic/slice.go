package main

import "fmt"

func main(){

    s1 := make([]int,3,8)
    fmt.Printf("cap: %d,len: %d\n",cap(s1),len(s1))
}

