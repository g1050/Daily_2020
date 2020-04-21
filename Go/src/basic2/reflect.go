package main

import "fmt"
import "reflect"

func main(){
    var x float64
    x = 3.14
    fmt.Println(reflect.TypeOf(x))
    fmt.Println(reflect.ValueOf(x))
    fmt.Println("Hello, World!")

    fmt.Println("----------------------")
    y := reflect.ValueOf(x)
    fmt.Println(y.Float())
}

