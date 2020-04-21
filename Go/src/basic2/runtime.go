package main

import "fmt"
import "runtime"

func main(){
    fmt.Println("Hello, World!")
    fmt.Println("GOROOT-->",runtime.GOROOT())
    fmt.Println("OS/platform",runtime.GOOS)
    fmt.Println("numcpu-->",runtime.NumCPU())
}

