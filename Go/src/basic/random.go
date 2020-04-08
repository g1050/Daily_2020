package main

import "fmt"
import "time"
import "math/rand"

func main(){

    for i:= 1;i<=100000000;i++{
        t := time.Now()
        rand.Seed(t.Unix())
        // sleep(1)
        fmt.Println(rand.Intn(10))
    }
}

