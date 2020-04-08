package main

import "fmt"

func main(){
    one := person{}
    one.name = "Jack"
    one.age = 22
    one.address = "Xi'an"
    fmt.Println(one)
}

type person struct{
    name string
    age int
    address string
}

