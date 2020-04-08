package main

import "fmt"

func main(){
    w1 := worker{name:"Jack",age:22}
    fmt.Println(w1)
    w1.work()
    w1.rest()

    w2 := &w1
    w2.rest()
}

type  worker struct{
    name string
    age int
}

func (w worker)work(){
    fmt.Println("He is working.")
}

func (w *worker)rest(){
    fmt.Println(w.name)
    fmt.Println("He is resting.")
}

